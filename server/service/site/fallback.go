package site

import (
	_ "embed"
	"strings"

	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func FallbackHandler(c *svcinfra.Context) {
	staticPath := config.Get("STATIC_PATH")
	if len(staticPath) != 0 {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/app") || strings.HasPrefix(path, "/login") {
			c.Context.File(staticPath + "/index.html")
		} else {
			c.Context.File(staticPath + path)
		}
	}
}
