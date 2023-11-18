package tmpl

import (
	"github.com/gin-gonic/gin"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func GetTmplTypeList(c *svcinfra.Context) {
	c.Bye(gin.H{"types": models.TmplTypeList()})
}
