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
	clustPath := "http://172.16.0.200:8080"
	baseURL := clustPath + "/api/namespaces/default/releases"
	// payload := `{"moudle":"codec","mode":"encoder","nmos_devname":"devname","2110-7_m_local_ip":"192.168.1.1","2110-7_b_local_ip":"192.168.1.2","input":{"in_g_2022-7":true,"in_v_protocol":"st2110-20","ipstream_master":{"in_v_src_address":"238.0.0.1:5004","in_v_dst_address":"239.0.0.1:5004","in_a_src_address":"238.0.0.1:5005","in_a_dst_address":"239.0.0.1:5005"},"ipstream_backup":{"in_v_src_address":"238.0.0.2:5004","in_v_dst_address":"239.0.0.2:5004","in_a_src_address":"238.0.0.2:5005","in_a_dst_address":"239.0.0.2:5005"},"videoformat":{"in_v_width":3840,"in_v_height":2160,"in_v_interlaced":false,"in_v_fps":50,"in_v_gamma":"hlg","in_v_gamut":"bt2020","in_v_compression_format":null,"in_v_compression_ratio":null},"audioformat":{"in_a_channels_number":16,"in_a_bits":24,"in_a_frequency":48000}},"output":{"out_g_2022-7":true,"out_g_local_ip1":"192.168.1.1","out_g_local_ip2":"192.168.1.2","out_params":{"out_v_protocol":"st2110-22","ipstream_master":{"out_v_src_address":"238.0.0.3:5004","out_v_dst_address":"192.168.1.1:5004","out_a_src_address":"238.0.0.3:5005","out_a_dst_address":"192.168.1.1:5005"},"ipstream_backup":{"out_v_src_address":"238.0.0.4:5004","out_v_dst_address":"192.168.1.2:5004","out_a_src_address":"238.0.0.4:5005","out_a_dst_address":"192.168.1.2:5005"},"videoformat":{"out_v_width":3840,"out_v_height":2160,"out_v_interlaced":false,"out_v_fps":50,"out_v_gamma":"hlg","out_v_gamut":"bt2020","out_v_compression_format":"jpeg-xs","out_v_compression_ratio":"8:1"},"audioformat":{"out_a_channels_number":16,"out_a_bits":24,"out_a_frequency":48000}}}}`

	// Create a new URL with the base URL
	apiURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	// Create query parameters
	// queryParams := url.Values{}
	// queryParams.Set("chart",, config.Get("CHART_PKG"))
	// // Add query parameters to the URL
	// apiURL.RawQuery = queryParams.Encode()

	// bodyBytes, err := json.Marshal(payload)
	// if err != nil {
	// 	fmt.Println("Error marshaling payload:", err)
	// 	return
	// }
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
