package cluster

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func buildProxyReq(c *svcinfra.Context, path string) *struct{} {
	clustPath := config.Get("CLUST_HOST")
	baseURL := clustPath + path
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}
	podReq, err := http.NewRequest("GET", apiURL.String(), nil)
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
	var resp2 struct{}
	json.Unmarshal(d, &resp2)
	return &resp2
}

func GetPodsPhase(c *svcinfra.Context) {
	var req struct {
		Device string `json:"device"`
	}
	c.ShouldBindJSON(&req)
	path := "/api/namespaces/default/releases/" + req.Device + "/podPhase"
	resp := buildProxyReq(c, path)
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
	resp := buildProxyReq(c, path)
	if resp == nil {
		c.GeneralError("删除pod失败")
		return
	}
	c.Bye(resp)
}

func GetNodes(c *svcinfra.Context) {
	path := "/api/nodes"
	resp := buildProxyReq(c, path)
	if resp == nil {
		c.GeneralError("获取节点列表失败")
		return
	}
	c.Bye(resp)
}

func GetNodeDetail(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	path := "/api/nodes/" + req.ID
	resp := buildProxyReq(c, path)
	if resp == nil {
		c.GeneralError("获取节点详情失败")
		return
	}
	c.Bye(resp)
}
