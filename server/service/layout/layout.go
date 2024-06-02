package layout

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/def"
	"vsoutlook.com/vsoutlook/infra/utils"
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
	deleteLayoutJson(layout)
	settings := syncLayoutToMvtSetting(layout, true)
	c.Bye(gin.H{"result": "ok", "settings": settings.Value})
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
	ly := models.QueryOne4[models.Layout]("name", req.Name, "deleted", 0)
	if ly != nil {
		c.GeneralError("布局名称已存在")
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
	settings := syncLayoutToMvtSetting(&newLayout, false)
	c.Bye(gin.H{"layout": newLayout.AsBasic(), "settings": settings.Value})
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
	if req.Name != layout.Name {
		ly := models.QueryOne4[models.Layout]("name", req.Name, "deleted", 0)
		if ly != nil {
			c.GeneralError("布局名称已存在")
			return
		}
	}
	if req.Size != def.LSize_4K && req.Size != def.LSize_HD {
		c.GeneralError("布局尺寸应为 4K 或 高清")
		return
	}
	layout.Name = req.Name
	layout.Size = req.Size
	models.Save(&layout)
	settings := syncLayoutToMvtSetting(layout, false)
	c.Bye(gin.H{"layout": layout.AsBasic(), "settings": settings.Value})
}

func Duplicate(c *svcinfra.Context) {
	var req struct {
		ID   string `json:"id"`
		Name string `json:"name"`
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
	ly := models.QueryOne4[models.Layout]("name", req.Name, "deleted", 0)
	if ly != nil {
		c.GeneralError("布局名称已存在")
		return
	}
	duplicated := models.Layout{
		Name:      req.Name,
		Size:      layout.Size,
		Content:   layout.Content,
		Published: 0,
	}
	models.Save(&duplicated)
	settings := syncLayoutToMvtSetting(&duplicated, false)
	c.Bye(gin.H{"layout": duplicated.AsBasic(), "settings": settings.Value})
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
	saveLayoutToJson(layout)

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
	settings := syncLayoutToMvtSetting(layout, false)
	c.Bye(gin.H{"result": "ok", "settings": settings.Value})
}

func saveLayoutToJson(layout *models.Layout) {
	path := config.AppDataDir + "/layouts"
	log.Printf("save layout to json: %s", path)
	utils.EnsureDirExists(path)
	filePath := fmt.Sprintf("%s/%s_%s.json", path, def.LSizeStr[layout.Size], layout.ID)
	err := utils.WriteJSONFile(filePath, layout.Content)
	if err != nil {
		log.Printf("failed to save layout to json: %v", err)
	}
}

func deleteLayoutJson(layout *models.Layout) {
	path := config.AppDataDir + "/layouts"
	filePath := fmt.Sprintf("%s/%s_%s.json", path, def.LSizeStr[layout.Size], layout.ID)
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("failed to delete layout json: %v", err)
	}
}

func syncLayoutToMvtSetting(layout *models.Layout, deleted bool) *models.Settings {
	settings := models.QuerySetting(def.SettingKey_Mtv)
	if layout.Published == 0 {
		return settings
	}
	var sVal string
	if settings == nil {
		sVal = "[]"
	} else {
		sVal = settings.Value
	}

	var mtvList []models.MtvSetting
	if err := json.Unmarshal([]byte(sVal), &mtvList); err != nil {
		log.Printf("sync layout to mtv setting err %s\n", err)
		return settings
	}
	var exist bool
	var deletedIndex int
	for i, mtv := range mtvList {
		if mtv.ID == layout.ID {
			if deleted {
				deletedIndex = i
			} else {
				mtvList[i].Name = layout.Name
				mtvList[i].Path = fmt.Sprintf("%s%s_%s.json", def.DefLayoutPath, def.LSizeStr[layout.Size], layout.ID)
			}
			exist = true
			break
		}
	}
	if !exist && !deleted {
		mtvList = append(mtvList, models.MtvSetting{
			ID:   layout.ID,
			Name: layout.Name,
			Path: fmt.Sprintf("%s%s_%s.json", def.DefLayoutPath, def.LSizeStr[layout.Size], layout.ID),
		})
	}
	if exist && deleted {
		mtvList = append(mtvList[:deletedIndex], mtvList[deletedIndex+1:]...)
	}
	newVal, err := json.Marshal(mtvList)
	if err != nil {
		log.Printf("sync layout to mtv setting err %s\n", err)
		return settings
	}
	if settings == nil {
		settings = &models.Settings{
			Key:   def.SettingKey_Mtv,
			Value: string(newVal),
		}
		models.Create(&settings)
	} else {
		settings.Value = string(newVal)
		models.Save(&settings)
	}
	return settings
}
