package tmpl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func GetTmplTypeList(c *svcinfra.Context) {
	c.Bye(gin.H{"types": models.TmplTypeList()})
}

func CreateTmplType(c *svcinfra.Context) {
	var req struct {
		Name     string `json:"name"`
		Icon     string `json:"icon"`
		Category string `json:"category"`
	}
	c.ShouldBindJSON(&req)
	if len(req.Name) == 0 {
		c.GeneralError("应用类型名称不能为空")
		return
	}
	if len(req.Icon) == 0 {
		c.GeneralError("应用类型图标不能为空")
		return
	}
	if len(req.Category) == 0 {
		c.GeneralError("应用类型类别不能为空")
		return
	}
	newTmpl := models.TmplType{
		Name:     req.Name,
		Icon:     req.Icon,
		Category: req.Category,
	}
	result := models.Create(&newTmpl)
	if result.Error != nil {
		fmt.Printf("failed to create tmpl in db: %v", result.Error)
		c.GeneralError("创建应用类型失败")
		return
	}
	c.Bye(gin.H{"tmpl": newTmpl.AsBasic()})
}
