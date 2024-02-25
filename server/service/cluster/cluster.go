package cluster

import (
	"bytes"
	"encoding/json"
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

type ClusterNodeDetail struct {
	NodeName    string                 `json:"nodeName"`
	NodeIP      string                 `json:"nodeIP"`
	Allocatable map[string]interface{} `json:"allocatable"`
}

func BuildProxyReq[T any](
	c *svcinfra.Context,
	method string,
	path string,
	query *map[string]interface{},
	params *map[string]interface{},
) *T {
	clustPath := config.Get("CLUST_HOST")
	baseURL := clustPath + path
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	if query != nil {
		// Create query parameters
		queryParams := url.Values{}
		fmt.Printf("query: %v", *query)
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
			return nil
		}
	}

	fmt.Println("Fetch URL:", apiURL.String())

	podReq, err := http.NewRequest(method, apiURL.String(), bytes.NewBuffer(paramBytes))
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return nil
	}
	podReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(podReq)
	if err != nil {
		fmt.Printf("failed to make request: %v", err)
		return nil
	}

	defer resp.Body.Close()

	d, err := io.ReadAll(resp.Body)
	fmt.Println("proxy node response:", string(d))
	if err != nil {
		fmt.Printf("failed to response: %v", err)
		return nil
	}

	var resp2 T

	err = json.Unmarshal(d, &resp2)
	if err != nil {
		d = []byte(fmt.Sprintf(`{"value": "%s"}`, string(d)))
		json.Unmarshal(d, &resp2)
	}

	return &resp2
}

func GetPodsPhase(c *svcinfra.Context) {
	var req struct {
		Device string `json:"device"`
	}
	c.ShouldBindJSON(&req)
	path := "/api/namespaces/default/releases/" + req.Device + "/podPhase"
	resp := BuildProxyReq[any](c, "GET", path, nil, nil)
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
	resp := BuildProxyReq[any](c, "GET", path, nil, nil)
	if resp == nil {
		c.GeneralError("删除pod失败")
		return
	}
	c.Bye(resp)
}

func GetNodes(c *svcinfra.Context) {
	path := "/nodes"
	resp := BuildProxyReq[[]interface{}](c, "GET", path, nil, nil)
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
	path := "/node"
	query := map[string]interface{}{
		"node_name": req.ID,
	}
	resp := BuildProxyReq[ClusterNodeDetail](c, "GET", path, &query, nil)
	if resp == nil {
		c.GeneralError("获取节点详情失败")
		return
	}
	data := make(map[string]interface{})
	data["node"] = resp

	node := models.ActiveNode(resp.NodeName)
	if node == nil {
		data["coreList"] = ""
	} else {
		data["coreList"] = node.CoreList
	}
	c.Bye(gin.H{"code": 0, "data": data})
}

func UpdateNode(c *svcinfra.Context) {
	var req struct {
		ID       string `json:"id"`
		CoreList string `json:"core"`
	}
	c.ShouldBindJSON(&req)
	node := models.ActiveNode(req.ID)
	if node == nil {
		node = &models.Node{
			ID:        req.ID,
			CoreList:  req.CoreList,
			Allocated: make(models.MapUint32Slice),
		}
	} else {
		if len(node.Allocated) == 0 {
			node.CoreList = req.CoreList
		} else {
			c.GeneralError("节点已分配，不可修改")
			return
		}
	}
	db.DB.Save(node)
	c.Bye(gin.H{"code": 0})
}
