package cluster

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/models/db"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

type ClusterListNode struct {
	NodeName string `json:"nodeName"`
	NodeIP   string `json:"nodeIP"`
}

type ClusterNodeDetail struct {
	NodeName     string                 `json:"nodeName"`
	NodeIP       string                 `json:"nodeIP"`
	Applications []string               `json:"applications"`
	Allocatable  map[string]interface{} `json:"allocatable"`
	Allocated    map[string]interface{} `json:"allocated"`
}

type ClusterNodeInfo struct {
	NodeName string `json:"nodeName"`
	NodeIP   string `json:"nodeIP"`
}

func BuildProxyReq[T any](
	c *svcinfra.Context,
	method string,
	path string,
	query *map[string]interface{},
	params *map[string]interface{},
	host string,
) (*T, error) {
	clustPath := config.Get("CLUST_HOST")
	if host != "" {
		clustPath = host
	}
	baseURL := clustPath + path
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, nil
	}

	if query != nil {
		// Create query parameters
		queryParams := url.Values{}
		fmt.Printf("query: %v\n", *query)
		for k, v := range *query {
			queryParams.Set(k, fmt.Sprintf("%v", v))
		}
		// Add query parameters to the URL
		apiURL.RawQuery = queryParams.Encode()
	}

	var paramBytes []byte
	if params != nil {
		paramBytes, err = json.Marshal(params)
		if err != nil {
			fmt.Printf("failed to stringify params: %v", err)
			return nil, nil
		} else {
			fmt.Printf("request params: %v\n", *params)
		}
	}

	fmt.Println("Fetch URL:", apiURL.String())

	podReq, err := http.NewRequest(method, apiURL.String(), bytes.NewBuffer(paramBytes))
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return nil, nil
	}
	podReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(podReq)
	if err != nil {
		fmt.Printf("failed to make request: %v", err)
		return nil, nil
	}

	defer resp.Body.Close()

	d, err := io.ReadAll(resp.Body)
	fmt.Println("proxy node response:", string(d))
	if err != nil {
		fmt.Printf("failed to response: %v", err)
		return nil, nil
	}

	var resp2 T

	err = json.Unmarshal(d, &resp2)
	if err != nil {
		return nil, errors.New(string(d))
	}

	return &resp2, nil
}

func GetPodsPhase(c *svcinfra.Context) {
	var req struct {
		Device string `json:"device"`
	}
	c.ShouldBindJSON(&req)
	path := "/api/namespaces/default/releases/" + req.Device + "/podPhase"
	resp, _ := BuildProxyReq[any](c, "GET", path, nil, nil, "")
	if resp == nil {
		c.GeneralError("获取pod状态失败")
		return
	}
	c.Bye(resp)
}

func DeletePod(c *svcinfra.Context) {
	var req struct {
		Pod string `json:"pod"`
	}
	c.ShouldBindJSON(&req)
	path := "/api/namespaces/default/pods/" + req.Pod + "/delete"
	resp, _ := BuildProxyReq[any](c, "GET", path, nil, nil, "")
	if resp == nil {
		c.GeneralError("删除pod失败")
		return
	}
	c.Bye(resp)
}

func GetNodes(c *svcinfra.Context) {
	path := "/nodes"
	resp, _ := BuildProxyReq[[]ClusterListNode](c, "GET", path, nil, nil, "")
	if resp == nil {
		c.GeneralError("获取节点列表失败")
		return
	}
	// Mock code
	// resp := make([]ClusterNodeInfo, 0)
	// resp = append(resp, ClusterNodeInfo{
	// 	NodeName: "controlplane",
	// 	NodeIP:   "192.168.1.12",
	// })
	// resp = append(resp, ClusterNodeInfo{
	// 	NodeName: "node4",
	// 	NodeIP:   "192.168.1.13",
	// })
	// end Mock code
	c.Bye(gin.H{"code": 0, "data": resp})
}

func GetNMosNodes(c *svcinfra.Context) {
	var req struct {
		Order string `json:"order"`
		Limit int    `json:"limit"`
	}
	c.ShouldBindJSON(&req)

	settings := models.GetSettings()

	path := "/x-nmos/query/v1.3/nodes"
	host := ""
	if rdsServerUrl, ok := settings["rds_server_url"]; !ok {
		c.GeneralError("RDS服务地址未设置")
		return
	} else {
		host = rdsServerUrl
	}

	if host == "" {
		c.GeneralError("NMOS_HOST not set")
		return
	}
	query := map[string]interface{}{
		"paging.order": req.Order,
		"paging.limit": req.Limit,
	}
	resp, _ := BuildProxyReq[[]any](c, "GET", path, &query, nil, host)
	if resp == nil {
		c.GeneralError("获取NMOS节点列表失败")
		return
	}
	c.Bye(gin.H{"code": 0, "data": resp})
}

func GetNodeDetail(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)

	data := make(map[string]interface{})

	// Mock code
	path := "/node"
	query := map[string]interface{}{
		"node_name": req.ID,
	}
	resp, _ := BuildProxyReq[ClusterNodeDetail](c, "GET", path, &query, nil, "")
	if resp != nil {
		data["node"] = resp
	}
	// end Mock code
	devices := models.GetDevicesByNode(req.ID)
	running := make([]string, 0)
	stopped := make([]string, 0)
	appMaps := make(map[string]bool)
	// Mock code
	for _, app := range resp.Applications {
		appMaps[app] = true
	}
	// end Mock code
	for _, device := range devices {
		if len(device.AppName) > 0 && appMaps[device.Name] {
			running = append(running, device.Name)
			continue
		}
		stopped = append(stopped, device.Name)
	}
	data["running"] = running
	data["stopped"] = stopped
	c.Bye(gin.H{"code": 0, "data": data})
}

func GetNodeNics(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	node := models.EnsureNode(req.ID)
	if node == nil {
		c.GeneralError("节点不存在")
		return
	}
	nics := node.Nics()
	result := make([]models.NicAsBasic, 0, len(nics))
	for _, t := range nics {
		result = append(result, t.AsBasic())
	}
	c.Bye(gin.H{"code": 0, "data": result})
}

func ReorderNodeNics(c *svcinfra.Context) {
	var req struct {
		ID       string `json:"id"`
		Position int64  `json:"position"`
	}
	c.ShouldBindJSON(&req)
	movedNic := models.ActiveNic(req.ID)
	if movedNic == nil {
		c.GeneralError("网卡不存在")
		return
	}
	tx := db.DB.Begin()
	oldPosition := movedNic.Position
	if req.Position == oldPosition {
		tx.Commit()
		c.Bye(gin.H{"code": 0})
		return
	}
	movedNic.Position = req.Position
	if req.Position > oldPosition {
		if err := tx.Model(&models.Nic{}).Where("node_id = ? and position > ? and position <= ?", movedNic.NodeID, oldPosition, req.Position).UpdateColumn("position", gorm.Expr("position - 1")).Error; err != nil {
			tx.Rollback()
			c.GeneralError("更新网卡位置失败")
			return
		}
	} else {
		if err := tx.Model(&models.Nic{}).Where("node_id = ? and position < ? and position >= ?", movedNic.NodeID, oldPosition, req.Position).UpdateColumn("position", gorm.Expr("position + 1")).Error; err != nil {
			tx.Rollback()
			c.GeneralError("更新网卡位置失败")
			return
		}
	}
	if err := tx.Save(movedNic).Error; err != nil {
		tx.Rollback()
		c.GeneralError("更新网卡位置失败")
		return
	}
	tx.Commit()
	var nics []models.Nic
	db.DB.Model(&models.Nic{}).Where("node_id=? and deleted=0", movedNic.NodeID).Order("position asc").Find(&nics)
	result := make([]models.NicAsBasic, 0, len(nics))
	for _, t := range nics {
		result = append(result, t.AsBasic())
	}
	c.Bye(gin.H{"code": 0, "data": result})
}

func CreateNic(c *svcinfra.Context) {
	var req struct {
		NodeID string `json:"nodeId"`
		models.NicInfo
	}
	c.ShouldBindJSON(&req)
	node := models.EnsureNode(req.NodeID)

	if len(req.CoreList) == 0 {
		c.GeneralError("网卡未设置核心")
		return
	}

	// Get the maximum position for the node's NICs
	var maxPosition int64
	db.DB.Model(&models.Nic{}).Where("node_id = ? AND deleted = 0", node.ID).Select("COALESCE(MAX(position), 0)").Scan(&maxPosition)

	nic := models.Nic{
		NodeID:        node.ID,
		NicNameMain:   req.NicNameMain,
		NicNameBackup: req.NicNameBackup,
		ReceiveMain:   req.ReceiveMain,
		ReceiveBackup: req.ReceiveBackup,
		SendMain:      req.SendMain,
		SendBackup:    req.SendBackup,
		CoreList:      req.CoreList,
		TxRxCoreList:  req.TxRxCoreList,
		DMAList:       req.DMAList,
		VFCount:       req.VFCount,
		Position:      maxPosition + 1, // Set position to be one more than the current maximum
	}
	db.DB.Create(&nic)
	c.Bye(gin.H{"code": 0, "data": nic.AsBasic()})
}

func UpdateNic(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
		models.NicInfo
	}
	c.ShouldBindJSON(&req)
	nic := models.ActiveNic(req.ID)
	if nic == nil {
		c.GeneralError("不存在")
		return
	}
	if len(req.CoreList) == 0 {
		c.GeneralError("网卡未设置核心")
		return
	}
	if len(nic.AllocatedCore) == 0 {
		nic.CoreList = req.CoreList
	} else if nic.CoreList != req.CoreList {
		c.GeneralError("核心已分配，不可修改")
		return
	}
	if len(nic.AllocatedDMA) == 0 {
		nic.DMAList = req.DMAList
	} else if nic.DMAList != req.DMAList {
		c.GeneralError("DMA通道已分配，不可修改")
		return
	}
	if len(nic.AllocatedCore) > int(req.VFCount) {
		c.GeneralError("VF数量不可小于已分配的数量")
		return
	}
	nic.NicNameMain = req.NicNameMain
	nic.NicNameBackup = req.NicNameBackup
	nic.ReceiveMain = req.ReceiveMain
	nic.ReceiveBackup = req.ReceiveBackup
	nic.SendMain = req.SendMain
	nic.SendBackup = req.SendBackup
	nic.CoreList = req.CoreList
	nic.TxRxCoreList = req.TxRxCoreList
	nic.DMAList = req.DMAList
	nic.VFCount = req.VFCount
	db.DB.Save(nic)
	c.Bye(gin.H{"code": 0, "data": nic.AsBasic()})
}

func DeleteNic(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	nic := models.ActiveNic(req.ID)
	if nic == nil {
		c.GeneralError("网卡不存在")
		return
	}
	if nic.IsDeleted() {
		c.GeneralError("网卡已删除")
		return
	}
	if len(nic.AllocatedCore) > 0 {
		c.GeneralError("网卡核心已分配，不可删除")
		return
	}
	if len(nic.AllocatedDMA) > 0 {
		c.GeneralError("网卡DMA通道已分配，不可删除")
		return
	}
	nic.Deleted = 1
	db.DB.Save(nic)
	c.Bye(gin.H{"code": 0})
}
