package device

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"vsoutlook.com/vsoutlook/infra/def"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/models/db"
	"vsoutlook.com/vsoutlook/service/cluster"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

type ClustResp struct {
	Code  int8        `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

type PodStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ClustReleaseApp struct {
	Name                string      `json:"name"`
	Namespace           string      `json:"namespace"`
	DisplayName         string      `json:"displayName"`
	ApplicationCategory string      `json:"applicationCategory"`
	TargetNode          string      `json:"targetNode"`
	PodsStatus          []PodStatus `json:"podsStatus,omitempty"`
}

type DeviceAsRelease struct {
	models.DeviceAsBasic
	TargetNode string      `json:"targetNode"`
	PodsStatus []PodStatus `json:"podsStatus,omitempty"`
}

type StringResp struct {
	Value string `json:"value"`
}

func GetDeviceList(c *svcinfra.Context) {
	clustDevices := make([]DeviceAsRelease, 0)
	clustQuery := map[string]interface{}{
		"release_status": "true",
	}
	resp2 := cluster.BuildProxyReq[[]ClustReleaseApp](c, "GET", "/apps", &clustQuery, nil)
	if resp2 == nil {
		fmt.Printf("failed to get device list")
		c.Bye(gin.H{"devices": clustDevices})
		return
	}

	for _, v := range *resp2 {
		device := models.QueryOne[models.Device]("app_name", v.Name)
		if device != nil {
			release := DeviceAsRelease{}
			basicDevice := device.AsBasic()
			copier.Copy(&release, &basicDevice)
			release.TargetNode = v.TargetNode
			release.PodsStatus = v.PodsStatus
			clustDevices = append(clustDevices, release)
		}
	}

	sort.Slice(clustDevices, func(i, j int) bool {
		return clustDevices[i].UpdatedAt > clustDevices[j].UpdatedAt
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
		ID   string `json:"id"`
		Node string `json:"node"`
	}
	c.ShouldBindJSON(&req)
	device := models.ActiveDevice(req.ID)
	if device == nil {
		c.Bye(gin.H{"result": "ok"})
		return
	}

	data := map[string]interface{}{
		"name": device.AppName,
	}
	resp2 := cluster.BuildProxyReq[struct {
		Name string `json:"name"`
	}](c, "POST", "/uninstall", nil, &data)
	if resp2 == nil {
		c.GeneralError("请求删除设备失败")
		return
	}

	if resp2.Name != device.AppName {
		c.GeneralError("删除设备失败")
		return
	}

	device.Deleted = def.SelfDeleted
	if !c.Save(&device) {
		return
	}

	node := models.ActiveNode(req.Node)
	if node != nil {
		delete(node.Allocated, device.ID)
		c.Save(&node)
	}

	c.Bye(gin.H{"result": "ok"})
}

func StartDevice(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	device := models.ActiveDevice(req.ID)
	if device == nil {
		c.GeneralError("启动的设备不存在")
		return
	}

	data := map[string]interface{}{
		"release_name":  device.AppName,
		"release_scale": "2",
	}
	resp2 := cluster.BuildProxyReq[StringResp](c, "GET", "/start", &data, nil)
	// start udx-i5urluof successfully with scale 2
	if resp2 == nil || !strings.Contains((*resp2).Value, "successfully") {
		c.GeneralError("请求启动设备失败")
		return
	}

	c.Bye(gin.H{"result": "ok"})
}

func StopDevice(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	device := models.ActiveDevice(req.ID)
	if device == nil {
		c.GeneralError("暂停的设备不存在")
		return
	}

	data := map[string]interface{}{
		"release_name": device.AppName,
	}
	resp2 := cluster.BuildProxyReq[StringResp](c, "GET", "/stop", &data, nil)
	fmt.Print(resp2)
	// stop udx-i5urluof successfully with scale 2
	if resp2 == nil || !strings.Contains((*resp2).Value, "successfully") {
		c.GeneralError("请求暂停设备失败")
		return
	}

	c.Bye(gin.H{"result": "ok"})
}

func allocateNodeCore(tmpl *models.Tmpl, node *models.Node) ([]uint32, error) {
	cpuNum := tmpl.Requirement.CpuNum
	if cpuNum <= 0 {
		return nil, errors.New("该应用没有设置cpu数量")
	}

	if len(node.CoreList) == 0 {
		return nil, errors.New("节点没有配置核心列表")
	}

	cores := utils.ParseCoreListString(node.CoreList)
	if len(cores) < int(cpuNum) {
		return nil, errors.New("节点核心数不足")
	}

	allocated := make(map[uint32]bool)
	for _, v := range node.Allocated {
		for _, v2 := range v {
			allocated[v2] = true
		}
	}

	if len(allocated)+cpuNum > len(cores) {
		return nil, errors.New("节点核心数不足")
	}

	var result []uint32
	for _, v := range cores {
		if !allocated[uint32(v)] {
			result = append(result, uint32(v))
		}
		if len(result) == int(cpuNum) {
			break
		}
	}

	return result, nil
}

func CreateDevice(c *svcinfra.Context) {
	var req struct {
		Name string `json:"name"`
		Tmpl string `json:"tmpl"`
		Node string `json:"node"`
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
	node := models.ActiveNode(req.Node)
	if node == nil {
		c.GeneralError("节点不存在")
		return
	}

	cores, err := allocateNodeCore(tmpl, node)
	if err != nil {
		c.GeneralError(err.Error())
		return
	}

	coreListStr := utils.IntArrayToString(cores)

	newDevice := models.Device{
		Name: req.Name,
		Tmpl: req.Tmpl,
	}

	tmplType := models.ActiveTmplType(tmpl.Type)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		fmt.Printf("parse request body: %v", err)
		c.GeneralError("请求数据格式错误")
		db.DB.Delete(&newDevice)
		return
	}
	data["displayName"] = req.Name
	data["applicationCategory"] = tmplType.Category
	data["coreList"] = coreListStr
	configFile := data["configFile"].(map[string]interface{})
	if _, ok := configFile["ipservice"]; ok {
		ipservice := configFile["ipservice"].(map[string]interface{})
		ipservice["binding_core_list"] = coreListStr
	}
	fmt.Printf("config file: %v", data["configFile"])
	newDevice.Config = utils.MapToString(configFile)

	resp2 := cluster.BuildProxyReq[struct {
		Name string `json:"name"`
	}](c, "POST", "/install", nil, &data)
	if resp2 == nil {
		fmt.Printf("failed to parse response: %v", err)
		c.GeneralError("部署节点返回数据格式错误")
		db.DB.Delete(&newDevice)
		return
	}
	newDevice.AppName = resp2.Name
	result := models.Create(&newDevice)
	if result.Error != nil {
		fmt.Printf("failed to create device in db: %v", result.Error)
		c.GeneralError("部署设备失败")
		return
	}

	node.Allocated[newDevice.ID] = cores
	c.Save(&node)

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

	c.Bye(gin.H{"result": "ok"})
}
