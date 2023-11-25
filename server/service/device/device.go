package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/def"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/models/db"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

type ClustResp struct {
	Code  int8        `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type ClustRelease struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Revision     string `json:"revision"`
	Updated      string `json:"updated"`
	Status       string `json:"status"`
	Chart        string `json:"chart"`
	ChartVersion string `json:"chart_version"`
	AppVersion   string `json:"app_version"`
}

type DeviceAsRelease struct {
	models.DeviceAsBasic
	NodeName     string `json:"nodeName"`
	Revision     string `json:"revision"`
	ChartVersion string `json:"chartVersion"`
	AppVersion   string `json:"appVersion"`
	Status       string `json:"status"`
	Updated      string `json:"updated"`
}

func GetDeviceList(c *svcinfra.Context) {
	var clustDevices []DeviceAsRelease
	clustPath := config.Get("CLUST_HOST")
	clustReq, err := http.NewRequest("GET", clustPath+"/api/namespaces/default/releases", nil)
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		c.Bye(gin.H{"devices": clustDevices})
		return
	}
	clustReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(clustReq)
	if err != nil {
		fmt.Printf("failed to make request: %v", err)
		c.Bye(gin.H{"devices": clustDevices})
		return
	}

	defer resp.Body.Close()

	d, _ := io.ReadAll(resp.Body)
	fmt.Println("cluster node response:", string(d))
	var resp2 ClustResp
	json.Unmarshal(d, &resp2)
	if resp2.Code != 0 {
		fmt.Printf("failed to make request: %v", resp2.Error)
		c.Bye(gin.H{"devices": clustDevices})
		return
	}

	data := resp2.Data.([]interface{})
	clustRelease := make([]ClustRelease, 0, len(data))
	for _, v := range data {
		release := ClustRelease{}
		decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:  &release,
			TagName: "json",
		})
		decoder.Decode(v)
		clustRelease = append(clustRelease, release)
	}
	releaseMap := make(map[string]ClustRelease)

	for _, release := range clustRelease {
		releaseMap[release.Name] = release
	}

	devices := models.DeviceList()
	for _, device := range devices {
		var dr DeviceAsRelease
		if release, ok := releaseMap[device.Name]; ok {
			fmt.Println("release:", release)
			copier.Copy(&dr, &device)
			dr.NodeName = release.Name
			dr.Revision = release.Revision
			dr.Status = release.Status
			dr.ChartVersion = release.ChartVersion
			dr.AppVersion = release.AppVersion
			dr.Updated = release.Updated
			clustDevices = append(clustDevices, dr)
		}
	}

	c.Bye(gin.H{"devices": clustDevices})
}

func GetDevice(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	device := models.GetDevice(req.ID)
	c.Bye(gin.H{"tmpl": device.AsBasic()})
}

func DeleteDevice(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	device := models.ActiveDevice(req.ID)
	if device == nil {
		c.Bye(gin.H{"result": "ok"})
		return
	}
	device.Deleted = def.SelfDeleted
	if !c.Save(&device) {
		return
	}
	c.Bye(gin.H{"result": "ok"})
}

func CreateDevice(c *svcinfra.Context) {
	var req struct {
		Name string `json:"name"`
		Tmpl string `json:"tmpl"`
		Body string `json:"body"`
	}

	c.ShouldBindJSON(&req)
	if len(req.Name) == 0 {
		c.GeneralError("应用名称不能为空")
		return
	}
	if len(req.Tmpl) == 0 {
		c.GeneralError("应用类型不存在")
		return
	}

	newDevice := models.Device{
		Name: req.Name,
		Tmpl: req.Tmpl,
	}

	clustPath := config.Get("CLUST_HOST")
	baseURL := clustPath + "/api/namespaces/default/releases/" + newDevice.Name

	// Create a new URL with the base URL
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	// Create query parameters
	queryParams := url.Values{}
	queryParams.Set("chart", "local/foo")
	// Add query parameters to the URL
	apiURL.RawQuery = queryParams.Encode()

	fmt.Println("URL:", apiURL.String())

	clustReq, err := http.NewRequest("POST", apiURL.String(), strings.NewReader(req.Body))
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		c.GeneralError("创建部署节点请求失败")
		db.DB.Delete(&newDevice)
		return
	}
	clustReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(clustReq)
	if err != nil {
		fmt.Printf("failed to make request: %v", err)
		c.GeneralError("创建部署节点请求失败")
		db.DB.Delete(&newDevice)
		return
	}

	defer resp.Body.Close()

	d, _ := io.ReadAll(resp.Body)
	fmt.Println("cluster node response:", string(d))
	var resp2 ClustResp
	err = json.Unmarshal(d, &resp2)
	if err != nil {
		fmt.Printf("failed to parse response: %v", err)
		c.GeneralError("部署节点返回数据格式错误")
		db.DB.Delete(&newDevice)
		return
	}
	if len(resp2.Error) > 0 {
		fmt.Printf("failed to create device: %v", resp2.Error)
		c.GeneralError("请求部署节点失败")
		db.DB.Delete(&newDevice)
		return
	}

	result := models.Create(&newDevice)
	if result.Error != nil {
		c.GeneralError("创建应用失败")
		return
	}

	c.Bye(gin.H{"device": newDevice.AsBasic()})
}
