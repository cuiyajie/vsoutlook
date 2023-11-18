package site

import (
	_ "embed"
	"path/filepath"

	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func FallbackHandler(c *svcinfra.Context) {
	fpath := filepath.Join(config.AppDir, c.Request.URL.Path[5:])
	if utils.FileExists(fpath) {
		c.Context.File(fpath)
		return
	}
}
