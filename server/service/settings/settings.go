package settings

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"vsoutlook.com/vsoutlook/infra/def"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

func GetSettings(c *svcinfra.Context) {
	c.Bye(gin.H{"settings": models.GetSettings()})
}

func UpdateSetting(c *svcinfra.Context) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	c.ShouldBindJSON(&req)
	setting := models.QuerySetting(req.Key)
	if setting == nil {
		setting = &models.Settings{
			Key:   req.Key,
			Value: req.Value,
		}
		models.Create(&setting)
	} else {
		setting.Value = req.Value
		models.Save(&setting)
	}
	if req.Key == def.SettingKey_Mtv {
		syncMtvSettingToLayout(req.Value)
	}
	c.Bye(gin.H{"settings": setting.AsBasic()})
}

func syncMtvSettingToLayout(jsonStr string) {
	var mtvList []models.MtvSetting
	if err := json.Unmarshal([]byte(jsonStr), &mtvList); err != nil {
		log.Printf("sync mtv setting to layout err %s\n", err)
		return
	}
	var mtvMap = make(map[string]bool)
	for _, mtv := range mtvList {
		if mtv.ID == "" {
			continue
		}
		mtvMap[mtv.ID] = true
	}
	layouts := models.PublishedLayoutList()
	for _, layout := range layouts {
		if _, ok := mtvMap[layout.ID]; !ok {
			layout.Published = 0
			models.Save(&layout)
		}
	}
}
