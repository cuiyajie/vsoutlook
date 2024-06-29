package device

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
type ClustReleaseAppDetail struct {
	Name                string                 `json:"name"`
	Namespace           string                 `json:"namespace"`
	DisplayName         string                 `json:"displayName"`
	ApplicationCategory string                 `json:"applicationCategory"`
	TargetNode          string                 `json:"targetNode"`
	PodsStatus          map[string]interface{} `json:"podsStatus,omitempty"`
}

type DeviceAsRelease struct {
	models.DeviceAsBasic
	TargetNode string      `json:"targetNode"`
	NodeIP     string      `json:"nodeIP"`
	PodsStatus []PodStatus `json:"podsStatus,omitempty"`
}

type AuthService struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func generateDeviceSeedID() string {
	// Generate a version 4 (random) UUID
	u := uuid.New()

	// Convert the UUID to a string and remove hyphens
	uuidString := strings.ReplaceAll(u.String(), "-", "")

	// Format the UUID string as "8 characters - 4 characters - 4 characters - 4 characters - 12 characters"
	formattedUUID := fmt.Sprintf("%s-%s-%s-%s-%s", uuidString[:8], uuidString[8:12], uuidString[12:16], uuidString[16:20], uuidString[20:])

	return formattedUUID
}

func GetDeviceList(c *svcinfra.Context) {
	devices := models.DeviceList()
	if len(devices) == 0 {
		c.Bye(gin.H{"devices": devices})
		return
	}

	clustDevices := make([]DeviceAsRelease, 0)
	clustQuery := map[string]interface{}{
		"release_status": "true",
	}
	resp2, _ := cluster.BuildProxyReq[[]ClustReleaseApp](c, "GET", "/apps", &clustQuery, nil)
	if resp2 == nil {
		fmt.Printf("failed to get device list")
		c.Bye(gin.H{"devices": devices})
		return
	}

	resp3, _ := cluster.BuildProxyReq[[]cluster.ClusterNodeInfo](c, "GET", "/nodes", nil, nil)
	nodesMap := make(map[string]cluster.ClusterNodeInfo)
	if resp3 != nil {
		for _, v := range *resp3 {
			nodesMap[v.NodeName] = v
		}
	}

	appsMap := make(map[string]ClustReleaseApp)
	for _, v := range *resp2 {
		appsMap[v.Name] = v
	}

	stoppedDevices := make([]string, 0)
	for _, d := range devices {
		release := DeviceAsRelease{}
		copier.Copy(&release, &d)
		if app, ok := appsMap[d.AppName]; ok {
			release.TargetNode = app.TargetNode
			release.PodsStatus = app.PodsStatus
		} else {
			release.TargetNode = d.Node
			release.AppName = ""
			if d.AppName != "" {
				stoppedDevices = append(stoppedDevices, d.ID)
			}
		}
		if node, ok := nodesMap[release.TargetNode]; ok {
			release.Node = node.NodeName
			release.NodeIP = node.NodeIP
		} else {
			release.Node = ""
			release.NodeIP = ""
		}
		clustDevices = append(clustDevices, release)
	}
	db.DB.Model(&models.Device{}).Where("id in (?)", stoppedDevices).Updates(map[string]interface{}{"app_name": ""})

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
	c.Bye(gin.H{"device": device.AsDetail()})
}

func checkAllocatedCore(c *svcinfra.Context, node *models.Node) {
	if node.Allocated == nil {
		node.Allocated = make(models.MapUint32Slice)
	} else {
		for deviceId := range node.Allocated {
			device := models.ActiveDevice(deviceId)
			if device == nil || device.Deleted == def.SelfDeleted || device.AppName == "" || device.Node != node.ID {
				delete(node.Allocated, deviceId)
			}
		}
	}
	c.Save(&node)
}

func checkAllocatedDMA(c *svcinfra.Context, node *models.Node) {
	if node.AllocatedDMA == nil {
		node.AllocatedDMA = make(models.MapStringSlice)
	} else {
		for deviceId := range node.AllocatedDMA {
			device := models.ActiveDevice(deviceId)
			if device == nil || device.Deleted == def.SelfDeleted || device.AppName == "" || device.Node != node.ID {
				delete(node.AllocatedDMA, deviceId)
			}
		}
	}
	c.Save(&node)
}

func preInstallation(c *svcinfra.Context, configStr string, tmpl *models.Tmpl, node *models.Node, device *models.Device) (*[]uint32, *[]string, error) {
	checkAllocatedCore(c, node)
	cores, err := allocateNodeCore(tmpl, node)
	if err != nil {
		return nil, nil, err
	}

	checkAllocatedDMA(c, node)
	dmas, err := allocateNodeDMA(tmpl, node)
	if err != nil {
		return nil, nil, err
	}

	settings := models.GetSettings()

	coreListStr := utils.IntArrayToString(cores)
	dmaListStr := strings.Join(dmas, ",")
	var configFile map[string]interface{}
	err = json.Unmarshal([]byte(configStr), &configFile)
	if err != nil {
		return nil, nil, errors.New("设备配置文件格式错误")
	}
	fmt.Printf("configFile: %v\n", configFile)

	tmplType := models.ActiveTmplType(tmpl.Type)
	data := make(map[string]interface{})

	if paddr, ok := configFile["2110-7_m_local_ip"].(string); ok {
		data["primaryVFAddress"] = paddr + "/24"
	} else {
		return nil, nil, errors.New("设备主网卡地址格式错误")
	}
	if saddr, ok := configFile["2110-7_b_local_ip"].(string); ok {
		data["secondaryVFAddress"] = saddr + "/24"
	} else {
		return nil, nil, errors.New("设备备网卡地址格式错误")
	}
	data["targetNode"] = node.ID
	data["cpu"] = tmpl.Requirement.CpuNum
	data["memory"] = tmpl.Requirement.Memory
	data["hugepages"] = tmpl.Requirement.HugePage
	data["shm"] = tmpl.Requirement.Shm
	data["logLevel"] = tmpl.Requirement.LogLevel
	data["MaxRateMbpsByCore"] = tmpl.Requirement.MaxRateMbpsByCore
	data["RXsessioncnt"] = tmpl.Requirement.ReceiveSessions
	data["displayName"] = device.Name
	data["applicationCategory"] = tmplType.Category
	data["coreList"] = coreListStr
	data["DMAList"] = dmaListStr
	if _, ok := configFile["ipservice"]; ok {
		ipservice := configFile["ipservice"].(map[string]interface{})
		ipservice["max_bandwidth_percore"] = tmpl.Requirement.MaxRateMbpsByCore
		ipservice["log_level"] = tmpl.Requirement.LogLevel
		ipservice["binding_core_list"] = coreListStr
		ipservice["dma_list"] = dmaListStr
		ipservice["receive_sessions"] = tmpl.Requirement.ReceiveSessions
	}
	if len(device.SeedID) == 0 {
		device.SeedID = generateDeviceSeedID()
	}
	if _, ok := configFile["nmos"]; ok {
		nmos := configFile["nmos"].(map[string]interface{})
		if externalAddress, ok := nmos["host_addresses"].(string); ok {
			if len(externalAddress) > 0 {
				data["externalAddress"] = externalAddress
				externalPorts := make([]int, 0)
				if httpPort, ok1 := nmos["http_port"].(float64); ok1 {
					externalPorts = append(externalPorts, int(httpPort))
				}
				if tallyPort, ok2 := configFile["tally_port"].(float64); ok2 {
					externalPorts = append(externalPorts, int(tallyPort))
				}
				if apiServerPort, ok3 := configFile["api_server_port"].(float64); ok3 {
					externalPorts = append(externalPorts, int(apiServerPort))
				}
				fmt.Printf("externalPorts: %v\n", externalPorts)
				if len(externalPorts) > 0 {
					data["externalPorts"] = externalPorts
				}
			}
		}
		nmos["seed_id"] = device.SeedID
		if rdsServerUrl, ok := settings["rds_server_url"]; ok {
			nmos["rds_server_url"] = rdsServerUrl
		}
	}
	if asJsonStr, ok := settings["authorization_services"]; ok {
		var authServices []AuthService
		err := json.Unmarshal([]byte(asJsonStr), &authServices)
		authServiceData := make([]map[string]interface{}, 0)
		if err == nil {
			for i, v := range authServices {
				authServiceData = append(authServiceData, map[string]interface{}{
					"index": i,
					"ip":    v.IP,
					"port":  v.Port,
				})
			}
		}
		configFile["authorization_service"] = authServiceData
	}
	configJson := utils.MapToString(configFile)
	data["configFile"] = configJson
	device.Config = configJson

	resp2, err := cluster.BuildProxyReq[struct {
		Name string `json:"name"`
	}](c, "POST", "/install", nil, &data)
	if resp2 == nil {
		fmt.Printf("failed to parse response: %v", err)
		if err != nil {
			return nil, nil, fmt.Errorf("部署设备失败: %s", err.Error())
		} else {
			return nil, nil, errors.New("部署节点返回数据格式错误")
		}
	}
	device.AppName = resp2.Name
	return &cores, &dmas, nil
}

func preUninstallation(c *svcinfra.Context, device *models.Device) error {
	data := map[string]interface{}{
		"name": device.AppName,
	}
	respQ, _ := cluster.BuildProxyReq[ClustReleaseAppDetail](c, "GET", "/app", &map[string]interface{}{
		"release_name": device.AppName,
	}, nil)
	if respQ != nil {
		resp2, err := cluster.BuildProxyReq[struct {
			Name string `json:"name"`
		}](c, "POST", "/uninstall", nil, &data)
		if resp2 == nil {
			if err != nil {
				return fmt.Errorf("请求删除设备失败: %s", err.Error())
			} else {
				return errors.New("请求删除设备失败")
			}
		}

		if resp2.Name != device.AppName {
			return errors.New("删除设备失败")
		}
	}
	return nil
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

	if device.Node == "" {
		c.GeneralError("设备没有分配节点")
		return
	}

	tmpl := models.ActiveTmpl(device.Tmpl)
	if tmpl == nil {
		c.GeneralError("设备类型不存在")
		return
	}
	node := models.ActiveNode(device.Node)
	if node == nil {
		node = &models.Node{
			ID: device.Node,
		}
	}

	cores, dmas, err := preInstallation(c, device.Config, tmpl, node, device)
	if err != nil {
		c.GeneralError(err.Error())
		return
	}
	device.UpdatedAt = time.Now()
	c.Save(&device)
	node.Allocated[device.ID] = *cores
	node.AllocatedDMA[device.ID] = *dmas
	c.Save(&node)
	c.Bye(gin.H{"device": device.AsBasic()})
}

// deprecated
func StartDevice2(c *svcinfra.Context) {
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
	resp2, err := cluster.BuildProxyReq[any](c, "GET", "/start", &data, nil)
	// start udx-i5urluof successfully with scale 2
	if (resp2 == nil && err == nil) || (err != nil && !strings.Contains(err.Error(), "successfully")) {
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
		c.GeneralError("停止的设备不存在")
		return
	}
	if device.AppName == "" {
		c.Bye(gin.H{"result": "ok"})
		return
	}
	err := preUninstallation(c, device)
	if err != nil {
		c.GeneralError(err.Error())
		return
	}
	device.AppName = ""
	node := models.ActiveNode(device.Node)
	if node != nil {
		delete(node.Allocated, device.ID)
		delete(node.AllocatedDMA, device.ID)
		c.Save(&node)
	}
	c.Save(&device)
	c.Bye(gin.H{"result": "ok"})
}

// deprecated
func StopDevice2(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	device := models.ActiveDevice(req.ID)
	if device == nil {
		c.GeneralError("停止的设备不存在")
		return
	}

	data := map[string]interface{}{
		"release_name": device.AppName,
	}
	resp2, err := cluster.BuildProxyReq[any](c, "GET", "/stop", &data, nil)
	// stop udx-i5urluof successfully with scale 2
	if (resp2 == nil && err == nil) || (err != nil && !strings.Contains(err.Error(), "successfully")) {
		c.GeneralError("请求停止设备失败")
		return
	}

	c.Bye(gin.H{"result": "ok"})
}

func allocateNodeCore(tmpl *models.Tmpl, node *models.Node) ([]uint32, error) {
	dpdkCpu := tmpl.Requirement.DPDKCpu
	if dpdkCpu <= 0 {
		return nil, errors.New("该应用没有设置cpu数量")
	}

	if len(node.CoreList) == 0 {
		return nil, errors.New("节点没有配置核心列表")
	}

	if len(node.Allocated) >= int(node.VFCount) {
		return nil, errors.New("节点VF数量已满")
	}

	cores := utils.ParseCoreListString(node.CoreList)
	if len(cores) < int(dpdkCpu) {
		return nil, errors.New("节点核心数不足")
	}

	allocated := make(map[uint32]bool)
	for _, v := range node.Allocated {
		for _, v2 := range v {
			allocated[v2] = true
		}
	}

	if len(allocated)+dpdkCpu > len(cores) {
		return nil, errors.New("节点核心数不足")
	}

	var result []uint32
	for _, v := range cores {
		if !allocated[uint32(v)] {
			result = append(result, uint32(v))
		}
		if len(result) == int(dpdkCpu) {
			break
		}
	}

	return result, nil
}

func allocateNodeDMA(tmpl *models.Tmpl, node *models.Node) ([]string, error) {
	dma := tmpl.Requirement.DMA

	if dma > 0 && len(node.DMAList) == 0 {
		return nil, errors.New("节点没有配置DMA通道列表")
	}

	dmas := strings.Split(node.DMAList, ",")
	if len(dmas) < dma {
		return nil, errors.New("节点DMA通道数量不足")
	}

	allocated := make(map[string]bool)
	for _, v := range node.AllocatedDMA {
		for _, v2 := range v {
			allocated[v2] = true
		}
	}

	if len(allocated)+dma > len(dmas) {
		return nil, errors.New("节点DMA通道数量不足")
	}

	var result []string

	if dma == 0 {
		return result, nil
	}

	for _, v := range dmas {
		if !allocated[v] {
			result = append(result, v)
		}
		if len(result) == dma {
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

	d := models.QueryOne4[models.Device]("name", req.Name, "deleted", 0)
	if d != nil {
		c.GeneralError("设备名称已存在")
		return
	}

	tmpl := models.ActiveTmpl(req.Tmpl)
	if tmpl == nil {
		c.GeneralError("设备类型不存在")
		return
	}
	node := models.ActiveNode(req.Node)
	if node == nil {
		node = &models.Node{
			ID: req.Node,
		}
	}

	newDevice := models.Device{
		Name: req.Name,
		Tmpl: req.Tmpl,
		Node: req.Node,
	}

	cores, dmas, err := preInstallation(c, req.Body, tmpl, node, &newDevice)
	if err != nil {
		c.GeneralError(err.Error())
		return
	}
	result := models.Create(&newDevice)
	if result.Error != nil {
		fmt.Printf("failed to create device in db: %v", result.Error)
		c.GeneralError("部署设备失败")
		return
	}
	node.Allocated[newDevice.ID] = *cores
	node.AllocatedDMA[newDevice.ID] = *dmas
	c.Save(&node)
	c.Bye(gin.H{"device": newDevice.AsDetail()})
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
	if device == nil {
		c.GeneralError("设备不存在")
		return
	}

	device.Config = req.Body
	c.Save(&device)

	c.Bye(gin.H{"device": device.AsDetail()})
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
	if device.AppName != "" {
		err := preUninstallation(c, device)
		if err != nil {
			c.GeneralError(err.Error())
			return
		}
	}

	node := models.ActiveNode(device.Node)
	if node != nil {
		delete(node.Allocated, device.ID)
		delete(node.AllocatedDMA, device.ID)
		c.Save(&node)
	}

	device.AppName = ""
	device.Deleted = def.SelfDeleted
	if !c.Save(&device) {
		return
	}

	c.Bye(gin.H{"result": "ok"})
}
