package device

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
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

type DpdkInterface struct {
	Index    float64 `json:"physicalNicNumber"`
	MainIP   string  `json:"activeIPAddress"`
	BackupIP string  `json:"backupIPAddress"`
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
	resp2, _ := cluster.BuildProxyReq[[]ClustReleaseApp](c, "GET", "/apps", &clustQuery, nil, "")
	if resp2 == nil {
		fmt.Printf("failed to get device list")
		c.Bye(gin.H{"devices": devices})
		return
	}

	resp3, _ := cluster.BuildProxyReq[[]cluster.ClusterNodeInfo](c, "GET", "/nodes", nil, nil, "")
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

func checkAllocatedCore(c *svcinfra.Context, nic *models.Nic) {
	if nic.AllocatedCore == nil {
		nic.AllocatedCore = make(models.MapUint32Slice)
	} else {
		for deviceId := range nic.AllocatedCore {
			device := models.ActiveDevice(deviceId)
			if device == nil || device.Deleted == def.SelfDeleted || device.AppName == "" || device.Node != nic.NodeID {
				delete(nic.AllocatedCore, deviceId)
			}
		}
	}
	c.Save(&nic)
}

func checkAllocatedTxRx(c *svcinfra.Context, nic *models.Nic) {
	if nic.AllocatedTxRx == nil {
		nic.AllocatedTxRx = make(models.MapUint32Slice)
	} else {
		for deviceId := range nic.AllocatedTxRx {
			device := models.ActiveDevice(deviceId)
			if device == nil || device.Deleted == def.SelfDeleted || device.AppName == "" || device.Node != nic.NodeID {
				delete(nic.AllocatedTxRx, deviceId)
			}
		}
	}
	c.Save(&nic)
}

func checkAllocatedDMA(c *svcinfra.Context, nic *models.Nic) {
	if nic.AllocatedDMA == nil {
		nic.AllocatedDMA = make(models.MapStringSlice)
	} else {
		for deviceId := range nic.AllocatedDMA {
			device := models.ActiveDevice(deviceId)
			if device == nil || device.Deleted == def.SelfDeleted || device.AppName == "" || device.Node != nic.NodeID {
				delete(nic.AllocatedDMA, deviceId)
			}
		}
	}
	c.Save(&nic)
}

func preInstallation(c *svcinfra.Context, configStr string, tmpl *models.Tmpl, node *models.Node, device *models.Device) (*map[string][]uint32, *map[string][]string, *map[string][]uint32, error) {
	var configFile map[string]interface{}
	err := json.Unmarshal([]byte(configStr), &configFile)
	if err != nil {
		return nil, nil, nil, errors.New("设备配置文件格式错误")
	}

	coreType := tmpl.Requirement.CoreType
	if coreType == "" {
		coreType = "txrx"
	}
	if coreType == "none" {
		updateNestedObjectKeys(configFile, []string{"v_core", "a_core"}, -1)
	}
	rxShare := tmpl.Requirement.RxShare
	txShare := tmpl.Requirement.TxShare
	if rxShare == 0 {
		rxShare = 1
	}
	if txShare == 0 {
		txShare = 1
	}

	// fmt.Printf("configFile before handle: %v\n", configFile)
	rxCount := make(map[string]int)
	txCount := make(map[string]int)
	realRxCount := make(map[string]int)
	realTxCount := make(map[string]int)
	if coreType != "none" {
		countNodeCoreRecursively(configFile, rxCount, txCount)
		// iterate rx count and divide the value by rxShare then get the ceil value
		for k, v := range rxCount {
			realRxCount[k] = int(math.Ceil(float64(v) / float64(rxShare)))
		}
		// iterate tx count and divide the value by txShare then get the ceil value
		for k, v := range txCount {
			realTxCount[k] = int(math.Ceil(float64(v) / float64(txShare)))
		}
	}

	var coreMap = make(map[string][]uint32)
	var sdkMap = make(map[string][]uint32)
	var dmaMap = make(map[string][]string)
	var rxtxCoreMap = make(map[string][]uint32)

	coreListStr := ""
	sdkListStr := ""
	dmaListStr := ""
	dpdkinterfaces := make([]DpdkInterface, 0)
	if _, ok := configFile["nic_list"]; ok {
		nicTempList, ok1 := configFile["nic_list"].([]interface{})
		if !ok1 {
			return nil, nil, nil, errors.New("设备网卡配置错误")
		}
		var nicList []map[string]interface{}
		for _, item := range nicTempList {
			if m, ok := item.(map[string]interface{}); ok {
				nicList = append(nicList, m)
			} else {
				return nil, nil, nil, errors.New("设备网卡配置错误")
			}
		}

		if len(nicList) != 0 {
			// 循环网卡，有 error 就返回错误
			for nidx, nicItem := range nicList {
				nicId := nicItem["id"].(string)
				nic := models.ActiveNic(nicId)
				if nic == nil {
					return nil, nil, nil, fmt.Errorf("设备配置的网卡 %d 不存在", nidx)
				}
				rxCoreCount := realRxCount[strconv.Itoa(nidx)]
				txCoreCount := realTxCount[strconv.Itoa(nidx)]
				var dpdkRxCount int
				var dpdkTxCount int
				var rxtxRxCount int
				var rxtxTxCount int

				if coreType == "dpdk" {
					dpdkRxCount = rxCoreCount
					dpdkTxCount = txCoreCount
				} else if coreType == "txrx" {
					rxtxRxCount = rxCoreCount
					rxtxTxCount = txCoreCount
				}

				checkAllocatedCore(c, nic)
				checkAllocatedTxRx(c, nic)

				relatedDevices := make(map[string]bool)
				for key := range nic.AllocatedCore {
					relatedDevices[key] = true
				}
				for key := range nic.AllocatedTxRx {
					relatedDevices[key] = true
				}
				relatedDeviceCount := len(relatedDevices)

				if relatedDeviceCount >= int(nic.VFCount) {
					return nil, nil, nil, fmt.Errorf("网卡 %d VF数量已满", nidx)
				}

				dpdkCores, sdkCores, rxCores, txCores, err := allocateNodeCore(tmpl, nic, nidx, dpdkRxCount, dpdkTxCount)
				if err != nil {
					return nil, nil, nil, err
				}
				coreMap[nic.ID] = dpdkCores
				sdkMap[nic.ID] = sdkCores

				rxCores2, txCores2, err := allocatedNodeRxTx(nic, nidx, rxtxRxCount, rxtxTxCount)
				if err != nil {
					return nil, nil, nil, err
				}
				rxCores = append(rxCores, rxCores2...)
				txCores = append(txCores, txCores2...)
				rxtxCoreMap[fmt.Sprintf("rx#%d", nidx)] = expandArray(rxCores, rxShare)
				rxtxCoreMap[fmt.Sprintf("tx#%d", nidx)] = expandArray(txCores, txShare)

				checkAllocatedDMA(c, nic)
				dmas, err := allocateNodeDMA(tmpl, nic, nidx)
				if err != nil {
					return nil, nil, nil, err
				}
				dmaMap[nic.ID] = dmas

				dpdkinterfaces = append(dpdkinterfaces, DpdkInterface{
					Index:    nicItem["nicIndex"].(float64),
					MainIP:   nicItem["2110-7_m_local_ip"].(string),
					BackupIP: nicItem["2110-7_b_local_ip"].(string),
				})

				delete(nicItem, "id")
				delete(nicItem, "nicIndex")
			}
			coreListStr = utils.FormatUint32Array(utils.MergeMapArrays(coreMap))
			sdkListStr = utils.FormatUint32Array(utils.MergeMapArrays(sdkMap))
			dmaListStr = strings.Join(utils.MergeMapArrays(dmaMap), ",")
		}
	}

	if coreType != "none" {
		updateMultipleKeysRecursive(configFile, []string{"v_core", "a_core"}, rxtxCoreMap)
	}

	if _, ok := configFile["videoformat_enum"]; ok {
		vfs, ok := configFile["videoformat_enum"].([]interface{})
		if ok {
			for _, vf := range vfs {
				if vfm, ok := vf.(map[string]interface{}); ok {
					if _, ok := vfm["fps"]; ok {
						if fps, ok := vfm["fps"].(float64); ok {
							if math.Trunc(fps) == fps {
								vfm["fps"] = fmt.Sprintf("%.1f", fps)
							}
						}
					}
				}
			}
		}
	}

	settings := models.GetSettings()

	tmplType := models.ActiveTmplType(tmpl.Type)
	data := make(map[string]interface{})

	vsok := utils.Map{}
	if vsokJson, ok := settings["basic_send_receive_stream_config"]; ok {
		// check vsokJson is valid json
		err := json.Unmarshal([]byte(vsokJson), &vsok)
		if err != nil {
			return nil, nil, nil, errors.New("基础收发流配置格式错误")
		}
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
	data["applicationCategory"] = tmplType.AppCategory
	data["sendAVFrameNodeCount"] = tmpl.Requirement.SendAVFrameNodeCount
	data["recvFrameCount"] = tmpl.Requirement.RecvframeCount
	data["imageTag"] = tmpl.Requirement.ImageTag
	data["coreList"] = coreListStr
	data["sdkCoreList"] = sdkListStr
	data["DMAList"] = dmaListStr
	data["dpdkinterfaces"] = dpdkinterfaces
	if _, ok := configFile["used_signal_type"]; ok {
		usedSignalType := configFile["used_signal_type"].(float64)
		if usedSignalType == 1 {
			data["disableDPDKInterfaceAutoAllocation"] = true
			data["hostNetwork"] = true
		} else {
			data["disableDPDKInterfaceAutoAllocation"] = false
			data["hostNetwork"] = false
		}
	}
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
	externalPorts := make([]int, 0)
	if _, ok := configFile["nmos"]; ok {
		nmos := configFile["nmos"].(map[string]interface{})
		if externalAddress, ok := nmos["host_addresses"].(string); ok {
			if len(externalAddress) > 0 {
				data["externalAddress"] = externalAddress
			}
		}
		if httpPort, ok1 := nmos["http_port"].(float64); ok1 {
			externalPorts = append(externalPorts, int(httpPort))
		}
		nmos["seed_id"] = device.SeedID
	}
	// Add service ports from api_params to externalPorts
	if apiParams, ok := configFile["api_params"]; ok && apiParams != nil {
		if apiParamsArray, ok := apiParams.([]interface{}); ok {
			for _, param := range apiParamsArray {
				if paramMap, ok := param.(map[string]interface{}); ok {
					if servicePort, ok := paramMap["service_port"]; ok {
						if portNum, ok := servicePort.(float64); ok {
							externalPorts = append(externalPorts, int(portNum))
						}
					}
				}
			}
		}
	}
	if len(externalPorts) > 0 {
		fmt.Printf("externalPorts: %v\n", externalPorts)
		data["externalPorts"] = externalPorts
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
	fmt.Printf("configFile after handle: %v\n", configJson)
	data["configFile"] = configJson

	// vsok
	// Set LogLevel for all nested objects in vsok
	updateNestedObjectKeys(vsok, []string{"LogLevel", "log_level"}, tmpl.Requirement.LogLevel)
	vsok["maxrateMbpsbycore"] = tmpl.Requirement.MaxRateMbpsByCore
	vsok["rx_sessioncnt"] = tmpl.Requirement.ReceiveSessions
	vsok["SendAVFrameNodeCount"] = tmpl.Requirement.SendAVFrameNodeCount
	vsok["recvframe_cnt"] = tmpl.Requirement.RecvframeCount
	vsok["recvframecount"] = tmpl.Requirement.RecvframeCount
	vsok["dpdk_corelist"] = coreListStr
	vsok["sdk_corelist"] = sdkListStr
	vsok["dmalist"] = dmaListStr
	data["vsok"] = vsok

	device.Config = configJson

	// return nil, nil, nil, errors.New("调试部署")

	// print request data in json format
	jsonStr, _ := json.Marshal(data)
	fmt.Printf("install request data: %s\n", jsonStr)

	resp2, err := cluster.BuildProxyReq[struct {
		Name string `json:"name"`
	}](c, "POST", "/install", nil, &data, "")
	if resp2 == nil {
		fmt.Printf("failed to parse response: %v", err)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("部署设备失败: %s", err.Error())
		} else {
			return nil, nil, nil, errors.New("部署节点返回数据格式错误")
		}
	}
	device.AppName = resp2.Name
	return &coreMap, &dmaMap, &rxtxCoreMap, nil
}

func preUninstallation(c *svcinfra.Context, device *models.Device) error {
	data := map[string]interface{}{
		"name": device.AppName,
	}
	respQ, _ := cluster.BuildProxyReq[ClustReleaseAppDetail](c, "GET", "/app", &map[string]interface{}{
		"release_name": device.AppName,
	}, nil, "")
	if respQ != nil {
		resp2, err := cluster.BuildProxyReq[struct {
			Name string `json:"name"`
		}](c, "POST", "/uninstall", nil, &data, "")
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
	node := models.EnsureNode(device.Node)

	cores, dmas, rxtxCoreMap, err := preInstallation(c, device.Config, tmpl, node, device)
	if err != nil {
		c.GeneralError(err.Error())
		return
	}
	device.UpdatedAt = time.Now()
	c.Save(&device)

	for nicID, core := range *cores {
		nic := models.ActiveNic(nicID)
		if nic != nil {
			nic.AllocatedCore[device.ID] = core
			c.Save(&nic)
		}
	}
	for nicID, core := range *rxtxCoreMap {
		nic := models.ActiveNic(nicID)
		if nic != nil {
			nic.AllocatedTxRx[device.ID] = core
			c.Save(&nic)
		}
	}
	for nicID, dma := range *dmas {
		nic := models.ActiveNic(nicID)
		if nic != nil {
			nic.AllocatedDMA[device.ID] = dma
			c.Save(&nic)
		}
	}

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
	resp2, err := cluster.BuildProxyReq[any](c, "GET", "/start", &data, nil, "")
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
		nics := node.Nics()
		for _, nic := range nics {
			modified := false
			if _, exists := nic.AllocatedCore[device.ID]; exists {
				delete(nic.AllocatedCore, device.ID)
				modified = true
			}
			if _, exists := nic.AllocatedDMA[device.ID]; exists {
				delete(nic.AllocatedDMA, device.ID)
				modified = true
			}
			if _, exists := nic.AllocatedTxRx[device.ID]; exists {
				delete(nic.AllocatedTxRx, device.ID)
				modified = true
			}
			if modified {
				c.Save(&nic)
			}
		}
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
	resp2, err := cluster.BuildProxyReq[any](c, "GET", "/stop", &data, nil, "")
	// stop udx-i5urluof successfully with scale 2
	if (resp2 == nil && err == nil) || (err != nil && !strings.Contains(err.Error(), "successfully")) {
		c.GeneralError("请求停止设备失败")
		return
	}

	c.Bye(gin.H{"result": "ok"})
}

func allocateNodeCore(tmpl *models.Tmpl, nic *models.Nic, nidx int, rxCoreCount int, txCoreCount int) ([]uint32, []uint32, []uint32, []uint32, error) {
	nicConfig := tmpl.Requirement.NicConfig
	if nicConfig == nil {
		return nil, nil, nil, nil, errors.New("该应用没有设置网卡配置")
	}
	if nidx >= len(nicConfig) {
		return nil, nil, nil, nil, fmt.Errorf("第 %d 块网卡没有配置", nidx)
	}
	iConfig := &nicConfig[nidx]
	dpdkCpu := iConfig.DPDKCpu
	sdkCpu := iConfig.SDKCpu

	if len(nic.CoreList) == 0 && dpdkCpu+sdkCpu <= 0 {
		return nil, nil, nil, nil, fmt.Errorf("第 %d 块网卡没有配置Dpdk核心列表", nidx)
	}

	cores := utils.ParseCoreListString(nic.CoreList)
	if len(cores) < int(dpdkCpu)+int(sdkCpu)+rxCoreCount+txCoreCount {
		return nil, nil, nil, nil, fmt.Errorf("第 %d 块网卡核心数不足", nidx)
	}

	allocatedCore := make(map[uint32]bool)
	for _, v := range nic.AllocatedCore {
		for _, v2 := range v {
			allocatedCore[v2] = true
		}
	}

	if len(allocatedCore)+dpdkCpu+sdkCpu+rxCoreCount+txCoreCount > len(cores) {
		return nil, nil, nil, nil, fmt.Errorf("第 %d 块网卡核心数不足", nidx)
	}

	var dpdkCore []uint32
	var sdkCore []uint32
	var rxCore []uint32
	var txCore []uint32
	for _, v := range cores {
		if !allocatedCore[uint32(v)] {
			if len(dpdkCore) < dpdkCpu {
				dpdkCore = append(dpdkCore, uint32(v))
			} else if len(sdkCore) < sdkCpu {
				sdkCore = append(sdkCore, uint32(v))
			} else if len(rxCore) < rxCoreCount {
				rxCore = append(rxCore, uint32(v))
			} else if len(txCore) < txCoreCount {
				txCore = append(txCore, uint32(v))
			}
		}
		if len(dpdkCore) == dpdkCpu && len(sdkCore) == sdkCpu && len(rxCore) == rxCoreCount && len(txCore) == txCoreCount {
			break
		}
	}
	return dpdkCore, sdkCore, rxCore, txCore, nil
}

func allocatedNodeRxTx(nic *models.Nic, nidx int, rxCoreCount int, txCoreCount int) ([]uint32, []uint32, error) {
	cores := utils.ParseCoreListString(nic.TxRxCoreList)
	if len(cores) < rxCoreCount+txCoreCount {
		return nil, nil, fmt.Errorf("第 %d 块网卡应用收发流核心数不足", nidx)
	}

	allocatedTxRx := make(map[uint32]bool)
	for _, v := range nic.AllocatedTxRx {
		for _, v2 := range v {
			allocatedTxRx[v2] = true
		}
	}

	if len(allocatedTxRx)+rxCoreCount+txCoreCount > len(cores) {
		return nil, nil, fmt.Errorf("第 %d 块网卡核心数不足", nidx)
	}

	var rxCore []uint32
	var txCore []uint32
	for _, v := range cores {
		if !allocatedTxRx[uint32(v)] {
			if len(rxCore) < rxCoreCount {
				rxCore = append(rxCore, uint32(v))
			} else if len(txCore) < txCoreCount {
				txCore = append(txCore, uint32(v))
			}
		}
		if len(rxCore) == rxCoreCount && len(txCore) == txCoreCount {
			break
		}
	}
	return rxCore, txCore, nil
}

func allocateNodeDMA(tmpl *models.Tmpl, nic *models.Nic, nidx int) ([]string, error) {
	nicConfig := tmpl.Requirement.NicConfig
	if nicConfig == nil {
		return nil, errors.New("该应用没有设置网卡配置")
	}
	if nidx >= len(nicConfig) {
		return nil, fmt.Errorf("第 %d 块网卡没有配置", nidx)
	}
	iConfig := &nicConfig[nidx]
	dma := iConfig.DMA

	if dma > 0 && len(nic.DMAList) == 0 {
		return nil, fmt.Errorf("网卡 %d 没有配置DMA列表", nidx)
	}

	dmas := strings.Split(nic.DMAList, ",")
	if len(dmas) < dma {
		return nil, fmt.Errorf("网卡 %d DMA通道数量不足", nidx)
	}

	allocated := make(map[string]bool)
	for _, v := range nic.AllocatedDMA {
		for _, v2 := range v {
			allocated[v2] = true
		}
	}

	if len(allocated)+dma > len(dmas) {
		return nil, fmt.Errorf("网卡 %d DMA通道数量不足", nidx)
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
	node := models.EnsureNode(req.Node)

	newDevice := models.Device{
		Name: req.Name,
		Tmpl: req.Tmpl,
		Node: req.Node,
	}

	cores, dmas, rxTxCoreMap, err := preInstallation(c, req.Body, tmpl, node, &newDevice)
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
	for nicID, core := range *cores {
		nic := models.ActiveNic(nicID)
		if nic != nil {
			nic.AllocatedCore[newDevice.ID] = core
			c.Save(&nic)
		}
	}
	for nicID, core := range *rxTxCoreMap {
		nic := models.ActiveNic(nicID)
		if nic != nil {
			nic.AllocatedTxRx[newDevice.ID] = core
			c.Save(&nic)
		}
	}
	for nicID, dma := range *dmas {
		nic := models.ActiveNic(nicID)
		if nic != nil {
			nic.AllocatedDMA[newDevice.ID] = dma
			c.Save(&nic)
		}
	}
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
		nics := node.Nics()
		for _, nic := range nics {
			modified := false
			if _, exists := nic.AllocatedCore[device.ID]; exists {
				delete(nic.AllocatedCore, device.ID)
				modified = true
			}
			if _, exists := nic.AllocatedDMA[device.ID]; exists {
				delete(nic.AllocatedDMA, device.ID)
				modified = true
			}
			if _, exists := nic.AllocatedTxRx[device.ID]; exists {
				delete(nic.AllocatedTxRx, device.ID)
				modified = true
			}
			if modified {
				c.Save(&nic)
			}
		}
	}

	device.AppName = ""
	device.Deleted = def.SelfDeleted
	if !c.Save(&device) {
		return
	}

	c.Bye(gin.H{"result": "ok"})
}

// updateNestedObjectKeys recursively traverses a nested map structure and updates all occurrences
// of any of the specified keys with the provided value.
// Parameters:
//   - m: The map to update
//   - targetKeys: A slice of keys to look for in the map and nested maps
//   - newValue: The value to set for all occurrences of any targetKey
func updateNestedObjectKeys(m map[string]interface{}, targetKeys []string, newValue interface{}) {
	for k, v := range m {
		// Check if the current key is in the target keys list
		for _, targetKey := range targetKeys {
			if k == targetKey {
				m[k] = newValue
				break // No need to check other target keys
			}
		}

		// If the value is a map, recursively update it
		if nestedMap, ok := v.(map[string]interface{}); ok {
			updateNestedObjectKeys(nestedMap, targetKeys, newValue)
		}

		// If the value is a slice, check each element for maps
		if slice, ok := v.([]interface{}); ok {
			for _, item := range slice {
				if nestedMap, ok := item.(map[string]interface{}); ok {
					updateNestedObjectKeys(nestedMap, targetKeys, newValue)
				}
			}
		}
	}
}

// countNodeCoreRecursively recursively traverses a nested map structure and counts the occurrences by its value
func countNodeCoreRecursively(m map[string]interface{}, rxCount map[string]int, txCount map[string]int) {
	for k, v := range m {
		if k == "v_core" {
			// check v if is a string, if so, split v with # and check the prefix is rx or tx, and get the second part
			if vStr, ok := v.(string); ok {
				if strings.HasPrefix(vStr, "rx#") {
					// replace rx# with empty string
					rxCount[vStr[3:]]++
				} else if strings.HasPrefix(vStr, "tx#") {
					txCount[vStr[3:]]++
				}
			}
		}
		if nestedMap, ok := v.(map[string]interface{}); ok {
			countNodeCoreRecursively(nestedMap, rxCount, txCount)
		}
		if slice, ok := v.([]interface{}); ok {
			for _, item := range slice {
				if nestedMap, ok := item.(map[string]interface{}); ok {
					countNodeCoreRecursively(nestedMap, rxCount, txCount)
				}
			}
		}
	}
}

// updateMultipleKeysRecursive is a helper function that performs the recursive traversal
// and updates multiple keys at the same level with the same value from the array
func updateMultipleKeysRecursive(m map[string]interface{}, targetKeys []string, mapValues map[string][]uint32) {
	valueKey := ""
	for _, targetKey := range targetKeys {
		if _, exists := m[targetKey]; exists {
			valueKey = m[targetKey].(string)
			break
		}
	}

	// If we found at least one target key and we have values left
	if len(valueKey) > 0 && len(mapValues[valueKey]) > 0 {
		// Get the current value to use for all keys at this level
		currentValue := mapValues[valueKey][0]

		// Update all target keys that exist at this level
		for _, targetKey := range targetKeys {
			if _, exists := m[targetKey]; exists {
				m[targetKey] = currentValue
			}
		}

		// Remove the first element from the array after updating all keys at this level
		mapValues[valueKey] = mapValues[valueKey][1:]
	}

	// Continue recursively processing nested maps
	for _, v := range m {
		// If the value is a map, recursively update it
		if nestedMap, ok := v.(map[string]interface{}); ok {
			updateMultipleKeysRecursive(nestedMap, targetKeys, mapValues)
		}

		// If the value is a slice, check each element for maps
		if slice, ok := v.([]interface{}); ok {
			for _, item := range slice {
				if nestedMap, ok := item.(map[string]interface{}); ok {
					updateMultipleKeysRecursive(nestedMap, targetKeys, mapValues)
				}
			}
		}
	}
}

func expandArray(original []uint32, factor int) []uint32 {
	if factor <= 0 {
		return original // Return original array if factor is invalid
	}

	result := make([]uint32, 0, len(original)*factor)

	for _, value := range original {
		for i := 0; i < factor; i++ {
			result = append(result, value)
		}
	}

	return result
}
