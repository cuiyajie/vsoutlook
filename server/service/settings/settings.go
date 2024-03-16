package settings

import (
	"github.com/gin-gonic/gin"
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
	c.Bye(gin.H{"settings": setting.AsBasic()})
}
