package endswt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"vsoutlook.com/vsoutlook/infra/def"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type EndSwtStatus struct {
	Target uint8 `json:"target"`
	Source uint8 `json:"source"`
	Status uint8 `json:"status"`
}

type EndSwtSwitch struct {
	Target uint8 `json:"target"`
	Source uint8 `json:"source"`
}

func BuildProxyReq[T any](
	c *svcinfra.Context,
	method string,
	path string,
	query *map[string]interface{},
	params *map[string]interface{},
) (*T, *ErrorResponse) {
	settings := models.QuerySetting(def.SettingKey_EndSwtApi)
	var sVal string
	if settings == nil || !utils.IsValidIPv4WithPort(settings.Value) {
		sVal = ""
	} else {
		sVal = settings.Value
	}
	if len(sVal) == 0 {
		fmt.Println("setting end switch api address is empty", settings)
		return nil, nil
	}

	baseURL := "http://" + settings.Value + path
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

	var errorResponse ErrorResponse
	err = json.Unmarshal(d, &errorResponse)
	if err == nil && errorResponse.Message != "" {
		fmt.Println("Error response:", errorResponse.Message)
		return nil, &errorResponse
	}

	var resp2 T
	err = json.Unmarshal(d, &resp2)
	if err != nil {
		return nil, nil
	}

	return &resp2, nil
}

func GetEndSwtTarget(c *svcinfra.Context) {
	path := "/apis/v1/getstatus/0"
	resp, err := BuildProxyReq[EndSwtStatus](c, "GET", path, nil, nil)
	if resp == nil {
		if err == nil {
			c.GeneralError("获取末级切换目标列表失败")
		} else {
			c.GeneralError(err.Message)
		}
		return
	}
	c.Bye(gin.H{"data": resp})
}

func SwitchTarget(c *svcinfra.Context) {
	var req struct {
		Target uint8 `json:"target"`
		Source uint8 `json:"source"`
	}
	c.ShouldBindJSON(&req)

	data := make(map[string]interface{})

	path := "/apis/v1/switch"
	data["target"] = req.Target
	data["source"] = req.Source
	resp, err := BuildProxyReq[EndSwtSwitch](c, "POST", path, nil, &data)
	if resp == nil {
		if err == nil {
			c.GeneralError("切换末级切换目标失败")
		} else {
			c.GeneralError(err.Message)
		}
		return
	}
	c.Bye(gin.H{"data": resp})
}
