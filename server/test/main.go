package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("starting...")
	clustPath := "http://172.16.1.102:8080"
	baseURL := clustPath + "/api/namespaces/default/releases"

	// Create a new URL with the base URL
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("api url:", apiURL.String())
	clustReq, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return
	}
	clustReq.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(clustReq)
	if err != nil {
		fmt.Printf("failed to make request: %v", err)
		return
	}

	defer resp.Body.Close()

	d, _ := io.ReadAll(resp.Body)
	fmt.Println("cluster node response:", string(d))

	var resp2 struct {
		Code  int8   `json:"code"`
		Error string `json:"error,omitempty"`
	}
	err = json.Unmarshal(d, &resp2)
	if err != nil {
		fmt.Printf("failed to parse response: %v", err)
		return
	}
	if len(resp2.Error) > 0 {
		fmt.Printf("failed to create device: %v", resp2.Error)
		return
	}

	fmt.Println("success")
}
