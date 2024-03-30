package cluster

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
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
) (*T, error) {
	clustPath := config.Get("CLUST_HOST")
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
	client := &http.Client{}
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
	resp, _ := BuildProxyReq[any](c, "GET", path, nil, nil)
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
	resp, _ := BuildProxyReq[any](c, "GET", path, nil, nil)
	if resp == nil {
		c.GeneralError("删除pod失败")
		return
	}
	c.Bye(resp)
}

func GetNodes(c *svcinfra.Context) {
	path := "/nodes"
	resp, _ := BuildProxyReq[[]ClusterListNode](c, "GET", path, nil, nil)
	if resp == nil {
		c.GeneralError("获取节点列表失败")
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

	path := "/node"
	query := map[string]interface{}{
		"node_name": req.ID,
	}
	resp, _ := BuildProxyReq[ClusterNodeDetail](c, "GET", path, &query, nil)
	var node *models.Node
	if resp != nil {
		data["node"] = resp
	}
	devices := models.GetDevicesByNode(req.ID)
	running := make([]string, 0)
	stopped := make([]string, 0)
	appMaps := make(map[string]bool)
	for _, app := range resp.Applications {
		appMaps[app] = true
	}
	for _, device := range devices {
		if len(device.AppName) > 0 && appMaps[device.Name] {
			running = append(running, device.Name)
			continue
		}
		stopped = append(stopped, device.Name)
	}
	data["running"] = running
	data["stopped"] = stopped
	node = models.ActiveNode(req.ID)
	if node == nil {
		data["coreList"] = ""
		data["dmaList"] = ""
		data["vfCount"] = 0
	} else {
		data["coreList"] = node.CoreList
		data["dmaList"] = node.DMAList
		data["vfCount"] = node.VFCount
	}
	c.Bye(gin.H{"code": 0, "data": data})
}

func UpdateNode(c *svcinfra.Context) {
	var req struct {
		ID       string `json:"id"`
		CoreList string `json:"core"`
		DMAList  string `json:"dma"`
		VFCount  uint32 `json:"vf"`
	}
	c.ShouldBindJSON(&req)
	node := models.ActiveNode(req.ID)
	if node == nil {
		node = &models.Node{
			ID:        req.ID,
			CoreList:  req.CoreList,
			Allocated: make(models.MapUint32Slice),
			DMAList:   req.DMAList,
			VFCount:   req.VFCount,
		}
	} else {
		if len(node.Allocated) == 0 {
			node.CoreList = req.CoreList
		} else if node.CoreList != req.CoreList {
			c.GeneralError("节点已分配，不可修改")
			return
		}
		if len(node.AllocatedDMA) == 0 {
			node.DMAList = req.DMAList
		} else if node.DMAList != req.DMAList {
			c.GeneralError("DMA通道已分配，不可修改")
			return
		}
		if len(node.Allocated) > int(req.VFCount) {
			c.GeneralError("VF数量不可小于已分配的数量")
			return
		}
		node.DMAList = req.DMAList
		node.VFCount = req.VFCount
	}
	db.DB.Save(node)
	c.Bye(gin.H{"code": 0})
}
