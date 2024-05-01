package layout

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"vsoutlook.com/vsoutlook/infra/def"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func GetLayoutList(c *svcinfra.Context) {
	c.Bye(gin.H{"layouts": models.LayoutList()})
}

func GetLayout(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	layout := models.GetLayout(req.ID)
	c.Bye(gin.H{"layout": layout.AsBasic()})
}

func DeleteLayout(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	layout := models.ActiveLayout(req.ID)
	if layout == nil {
		c.Bye(gin.H{"result": "ok"})
		return
	}
	layout.Deleted = def.SelfDeleted
	if !c.Save(&layout) {
		return
	}
	c.Bye(gin.H{"result": "ok"})
}

func CreateLayout(c *svcinfra.Context) {
	var req struct {
		Name string `json:"name"`
		Size uint8  `json:"size"`
	}
	c.ShouldBindJSON(&req)
	if len(req.Name) == 0 {
		c.GeneralError("布局名称不能为空")
		return
	}
	if req.Size != def.LSize_4K && req.Size != def.LSize_HD {
		c.GeneralError("布局尺寸应为 4K 或 高清")
		return
	}
	newLayout := models.Layout{
		Name:      req.Name,
		Size:      req.Size,
		Published: 0,
	}
	result := models.Create(&newLayout)
	if result.Error != nil {
		fmt.Printf("failed to create tmpl in db: %v", result.Error)
		c.GeneralError("创建布局失败")
		return
	}
	c.Bye(gin.H{"layout": newLayout.AsBasic()})
}

func UpdateLayout(c *svcinfra.Context) {
	var req struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Size uint8  `json:"size"`
	}
	c.ShouldBindJSON(&req)
	layout := models.GetLayout(req.ID)
	if layout == nil || layout.Deleted != def.NotDeleted {
		c.GeneralError("布局不存在")
		return
	}
	if len(req.Name) == 0 {
		c.GeneralError("布局名称不能为空")
		return
	}
	if req.Size != def.LSize_4K && req.Size != def.LSize_HD {
		c.GeneralError("布局尺寸应为 4K 或 高清")
		return
	}
	layout.Name = req.Name
	layout.Size = req.Size
	models.Save(&layout)
	c.Bye(gin.H{"layout": layout.AsBasic()})
}

func UpdateLayoutLocation(c *svcinfra.Context) {
	var req struct {
		ID       string `json:"id"`
		Location string `json:"location"`
	}
	c.ShouldBindJSON(&req)
	layout := models.GetLayout(req.ID)
	if layout == nil || layout.Deleted != def.NotDeleted {
		c.GeneralError("布局不存在")
		return
	}
	if len(req.Location) == 0 {
		c.GeneralError("存储位置不能为空")
		return
	}
	// TODO: call cluster api to update layout location
	layout.Location = req.Location
	models.Save(&layout)
	c.Bye(gin.H{"layout": layout.AsBasic()})
}

func UpdateLayoutContent(c *svcinfra.Context) {
	var req struct {
		ID      string                           `json:"id"`
		Content datatypes.JSONSlice[interface{}] `json:"content"`
	}
	c.ShouldBindJSON(&req)
	layout := models.GetLayout(req.ID)
	if layout == nil || layout.Deleted != def.NotDeleted {
		c.GeneralError("布局不存在")
		return
	}
	layout.Content = req.Content
	models.Save(&layout)
	c.Bye(gin.H{"layout": layout.AsBasic()})
}

func PublishLayout(c *svcinfra.Context) {
	var req struct {
		ID string `json:"id"`
	}
	c.ShouldBindJSON(&req)
	layout := models.GetLayout(req.ID)
	if layout == nil || layout.Deleted != def.NotDeleted {
		c.GeneralError("布局不存在")
		return
	}
	layout.Published = 1
	models.Save(&layout)
	c.Bye(gin.H{"result": "ok"})
}
