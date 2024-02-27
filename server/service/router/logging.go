package router

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/track"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func logger(c *gin.Context) {
	if config.Testing {
		c.Next()
		return
	}

	if c.Request.Method != http.MethodPost ||
		c.Request.ContentLength > 100_000 ||
		!strings.HasPrefix(c.Request.URL.Path, "/api") {
		c.Next()
		return
	}

	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := io.ReadAll(tee)
	c.Request.Body = io.NopCloser(&buf)

	var info utils.Map
	ctype := c.Request.Header["Content-Type"]
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/api") && len(ctype) == 1 && strings.Contains(ctype[0], "json") {
		if err := json.Unmarshal(body, &info); err != nil {
			log.Printf("unmarshal body err: %s body: %s", err, body)
		}
	}

	data := string(body)
	if len(data) > 1000 {
		data = data[:1000]
	}

	if strings.HasPrefix(path, "api") && !strings.Contains(data, `"password"`) {
		log.Println("body raw:", data)
		c.Set("bodyPreview", data)
	}
	c.Next()

	if info == nil {
		return
	}

	// c.Request.RequestURI is empty in test env
	info[track.Path] = c.Request.URL.Path
	info[track.UserID] = c.GetString(svcinfra.UIDKey)
	info[track.EventType] = "api"
	info[track.UserAgent] = c.Request.UserAgent()

	if params, ok := c.Get(svcinfra.CtxReqParams); ok {
		info[track.ReqParams] = params
	}

	sigBundle := strings.Split(c.Request.Header.Get("X-Signature"), ",")
	// todo: remove len(sigBundle) == 4 after all clients are updated
	if len(sigBundle) == 4 || len(sigBundle) == 5 {
		info[track.BrowserID] = sigBundle[2]
	}
	track.Log(info)
}
