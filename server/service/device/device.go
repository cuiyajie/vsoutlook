package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
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
	Error string      `json:"error,omitempty"`
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
	Phase        string `json:"phase"`
}

type DeviceAsRelease struct {
	models.DeviceAsBasic
	Release      string `json:"release"`
	Revision     string `json:"revision"`
	ChartVersion string `json:"chartVersion"`
	AppVersion   string `json:"appVersion"`
	Status       string `json:"status"`
	Updated      string `json:"updated"`
	Phase        string `json:"phase"`
}

func GetDeviceList(c *svcinfra.Context) {
	var clustDevices []DeviceAsRelease
	clustPath := config.Get("CLUST_HOST")
	baseURL := clustPath + "/api/namespaces/default/releases"
	// Create a new URL with the base URL
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	// Create query parameters
	queryParams := url.Values{}
	queryParams.Set("release_status", "true")
	// Add query parameters to the URL
	apiURL.RawQuery = queryParams.Encode()

	fmt.Println("Fetch URL:", apiURL.String())

	clustReq, err := http.NewRequest("GET", apiURL.String(), nil)
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
			dr.Release = release.Name
			dr.Revision = release.Revision
			dr.Status = release.Status
			dr.ChartVersion = release.ChartVersion
			dr.AppVersion = release.AppVersion
			dr.Updated = release.Updated
			dr.Phase = release.Phase
			clustDevices = append(clustDevices, dr)
		}
	}

	sort.Slice(clustDevices, func(i, j int) bool {
		return clustDevices[i].Updated > clustDevices[j].Updated
	})

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

	clustPath := config.Get("CLUST_HOST")
	baseUrl := clustPath + "/api/namespaces/default/releases/" + device.Name
	fmt.Println("Delete URL:", baseUrl)
	clustReq, err := http.NewRequest("DELETE", baseUrl, nil)
	if err != nil {
		fmt.Printf("failed to delete request: %v", err)
		c.GeneralError("创建删除请求失败")
		return
	}
	client := &http.Client{}
	resp, err := client.Do(clustReq)
	if err != nil {
		fmt.Printf("failed to make request: %v", err)
		c.GeneralError("发送删除请求失败")
		return
	}

	defer resp.Body.Close()

	d, _ := io.ReadAll(resp.Body)
	fmt.Println("cluster node response:", string(d))
	var resp2 ClustResp
	err = json.Unmarshal(d, &resp2)
	if err != nil {
		fmt.Printf("failed to parse request: %v", err)
		c.GeneralError("请求删除设备失败")
		return
	}
	if resp2.Code != 0 && len(resp2.Error) > 0 && !strings.HasSuffix(resp2.Error, "not found") {
		fmt.Printf("failed to make request: %v", resp2.Error)
		c.GeneralError("删除设备失败")
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

	tmpl := models.ActiveTmpl(req.Tmpl)
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
	chart := config.Get("CHART_PKG")
	if len(tmpl.Requirement.Chart) > 0 {
		chart = tmpl.Requirement.Chart
	}
	queryParams.Set("chart", chart)
	// Add query parameters to the URL
	apiURL.RawQuery = queryParams.Encode()

	fmt.Println("Create URL:", apiURL.String(), req.Body)
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
		c.GeneralError("部署设备失败")
		return
	}

	c.Bye(gin.H{"device": newDevice.AsBasic()})
}

func UpdateDevice(c *svcinfra.Context) {
	var req struct {
		Name     string `json:"name"`
		DeviceId string `json:"deviceId"`
		Body     string `json:"body"`
	}

	c.ShouldBindJSON(&req)
	if len(req.Name) == 0 {
		c.GeneralError("设备名称不能为空")
		return
	}

	device := models.ActiveDevice(req.DeviceId)
	tmpl := models.ActiveTmpl(device.Tmpl)

	if req.Name != device.Name {
		c.GeneralError("设备不存在")
		return
	}

	clustPath := config.Get("CLUST_HOST")
	baseURL := clustPath + "/api/namespaces/default/releases/" + device.Name

	// Create a new URL with the base URL
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	// Create query parameters
	queryParams := url.Values{}
	chart := config.Get("CHART_PKG")
	if len(tmpl.Requirement.Chart) > 0 {
		chart = tmpl.Requirement.Chart
	}
	queryParams.Set("chart", chart)
	// Add query parameters to the URL
	apiURL.RawQuery = queryParams.Encode()

	fmt.Println("Update URL:", apiURL.String())

	clustReq, err := http.NewRequest("PUT", apiURL.String(), strings.NewReader(req.Body))
	if err != nil {
		fmt.Printf("failed to create update request: %v", err)
		c.GeneralError("创建更新设备请求失败")
		return
	}
	clustReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(clustReq)
	if err != nil {
		fmt.Printf("failed to make update request: %v", err)
		c.GeneralError("发送更新设备请求失败")
		return
	}

	defer resp.Body.Close()

	d, _ := io.ReadAll(resp.Body)
	fmt.Println("cluster node response:", string(d))
	var resp2 ClustResp
	err = json.Unmarshal(d, &resp2)
	if err != nil {
		fmt.Printf("failed to parse response: %v", err)
		c.GeneralError("更新设备返回数据格式错误")
		return
	}
	if len(resp2.Error) > 0 {
		fmt.Printf("failed to update device: %v", resp2.Error)
		c.GeneralError("更新设备失败")
		return
	}

	models.Save(&device)
	c.Bye(gin.H{"device": device.AsBasic()})
}
