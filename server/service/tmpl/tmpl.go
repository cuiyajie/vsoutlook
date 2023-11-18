package tmpl

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"vsoutlook.com/vsoutlook/infra/def"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func GetTmplList(c *svcinfra.Context) {
	c.Bye(gin.H{"tmpls": models.TmplList()})
}

func GetTmpl(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	tmpl := models.GetTmpl(req.ID)
	c.Bye(gin.H{"tmpl": tmpl.AsDetail()})
}

func DeleteTmpl(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	tmpl := models.ActiveTmpl(req.ID)
	if tmpl == nil {
		c.Bye(gin.H{"result": "ok"})
		return
	}
	tmpl.Deleted = def.SelfDeleted
	if !c.Save(&tmpl) {
		return
	}
	c.Bye(gin.H{"result": "ok"})
}

func CreateTmpl(c *svcinfra.Context) {
	var req struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	c.ShouldBindJSON(&req)
	if len(req.Name) == 0 {
		c.GeneralError("应用名称不能为空")
		return
	}
	if len(req.Type) == 0 {
		c.GeneralError("应用类型不存在")
		return
	}
	tmplType := models.ActiveTmplType(req.Type)
	if tmplType == nil {
		c.GeneralError("应用类型不存在")
		return
	}
	newTmpl := models.Tmpl{
		Name: req.Name,
		Type: tmplType.ID,
	}
	result := models.Create(&newTmpl)
	if result.Error != nil {
		c.GeneralError("创建应用失败")
		return
	}
	c.Bye(gin.H{"tmpl": newTmpl.AsBasic()})
}

func UpdateTmpl(c *svcinfra.Context) {
	var req struct {
		ID          string    `json:"id"`
		Flow        string    `json:"flow"`
		Requirement utils.Map `json:"requirement"`
		Description string    `json:"description"`
	}
	c.ShouldBindJSON(&req)
	tmpl := models.GetTmpl(req.ID)
	if tmpl == nil || tmpl.Deleted != def.NotDeleted {
		c.GeneralError("应用不存在")
		return
	}
	tmpl.Flow = req.Flow
	rq := models.TmplRequirement{}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &rq,
		TagName: "json",
	})
	if err != nil {
		c.GeneralError("应用配置错误")
		return
	}
	err = decoder.Decode(req.Requirement)
	if err != nil {
		c.GeneralError("应用配置错误")
		return
	}
	tmpl.Requirement = rq
	tmpl.Description = req.Description
	models.Save(&tmpl)
	c.Bye(gin.H{"tmpl": tmpl.AsDetail()})
}

func SetTmplListed(c *svcinfra.Context) {
	var req struct {
		ID     string `json:"id"`
		Listed bool   `json:"listed"`
	}
	c.ShouldBindJSON(&req)
	tmpl := models.GetTmpl(req.ID)
	if tmpl == nil || tmpl.Deleted != def.NotDeleted {
		c.GeneralError("应用不存在")
		return
	}
	if req.Listed {
		tmpl.Listed = 1
	} else {
		tmpl.Listed = 0
	}
	models.Save(&tmpl)
	c.Bye(gin.H{"result": "ok"})
}
