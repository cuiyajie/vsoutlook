package backup

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/models/db"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

// export all data from database to a sql file
func ExportData(c *svcinfra.Context) {
	var req struct {
		Tables []string `json:"tables"`
	}
	c.ShouldBindJSON(&req)
	data := gin.H{}
	if len(req.Tables) == 0 {
		c.Bye(gin.H{"data": data})
		return
	}
	if utils.StringContains(req.Tables, "tmpl_types") {
		tmplTypes := models.TmplTypeList()
		data["tmpl_types"] = make([]gin.H, 0, len(tmplTypes))
		for _, tt := range tmplTypes {
			data["tmpl_types"] = append(data["tmpl_types"].([]gin.H), gin.H{
				"name":     tt.Name,
				"icon":     tt.Icon,
				"category": tt.Category,
			})
		}
	}
	if utils.StringContains(req.Tables, "tmpls") {
		var tmpls []models.Tmpl
		db.DB.Where("deleted=0").Find(&tmpls)
		data["tmpls"] = make([]gin.H, 0, len(tmpls))
		for _, tmpl := range tmpls {
			tmplType := models.TmplTypeBy("id", tmpl.Type)
			data["tmpls"] = append(data["tmpls"].([]gin.H), gin.H{
				"name":        tmpl.Name,
				"type":        tmplType.Category,
				"description": tmpl.Description,
				"requirement": tmpl.Requirement,
				"flow":        tmpl.Flow,
				"listed":      tmpl.Listed,
			})
		}
	}
	if utils.StringContains(req.Tables, "devices") {
		var devices []models.Device
		db.DB.Where("deleted=0").Find(&devices)
		data["devices"] = make([]gin.H, 0, len(devices))
		for _, device := range devices {
			deviceBasic := device.AsDetail()
			data["devices"] = append(data["devices"].([]gin.H), gin.H{
				"name":      device.Name,
				"tmpl_name": deviceBasic.TmplName,
				"tmpl_type": deviceBasic.TmplTypeCategory,
				"config":    deviceBasic.Config,
				"node":      deviceBasic.Node,
			})
		}
	}
	if utils.StringContains(req.Tables, "layouts") {
		var layouts []models.Layout
		db.DB.Where("deleted=0").Find(&layouts)
		data["layouts"] = make([]gin.H, 0, len(layouts))
		for _, layout := range layouts {
			data["layouts"] = append(data["layouts"].([]gin.H), gin.H{
				"name":    layout.Name,
				"size":    layout.Size,
				"publish": layout.Published,
				"content": layout.Content,
			})
		}
	}
	if utils.StringContains(req.Tables, "settings") {
		var settings []models.Settings
		db.DB.Find(&settings)
		data["settings"] = make([]gin.H, 0, len(settings))
		for _, setting := range settings {
			data["settings"] = append(data["settings"].([]gin.H), gin.H{
				"key":   setting.Key,
				"value": setting.Value,
			})
		}
	}
	c.Bye(gin.H{"data": data})
}

// import data from a sql file to database
func ImportData(c *svcinfra.Context) {
	var req struct {
		File string `json:"file"`
		Skip bool   `json:"skip"`
	}
	c.ShouldBindJSON(&req)
	if req.File == "" {
		c.GeneralError("请上传正确的数据文件")
		return
	}
	var data struct {
		TmplTypes []struct {
			Name     string `json:"name"`
			Icon     string `json:"icon"`
			Category string `json:"category"`
		} `json:"tmpl_types"`
		Tmpls []struct {
			Name        string                 `json:"name"`
			Type        string                 `json:"type"`
			Description string                 `json:"description"`
			Requirement models.TmplRequirement `json:"requirement"`
			Flow        string                 `json:"flow"`
			Listed      uint8                  `json:"listed"`
		} `json:"tmpls"`
		Devices []struct {
			Name     string `json:"name"`
			Tmpl     string `json:"tmpl_name"`
			TmplType string `json:"tmpl_type"`
			Config   string `json:"config"`
			Node     string `json:"node"`
		} `json:"devices"`
		Layouts []struct {
			Name    string                           `json:"name"`
			Size    uint8                            `json:"size"`
			Publish uint8                            `json:"publish"`
			Content datatypes.JSONSlice[interface{}] `json:"content"`
		} `json:"layouts"`
		Settings []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"settings"`
	}
	err := json.Unmarshal([]byte(req.File), &data)
	if err != nil {
		c.GeneralError("解析数据文件失败，请上传正确的数据文件")
		return
	}

	for _, tt := range data.TmplTypes {
		tmplType := models.TmplTypeBy("category", tt.Category)
		if tmplType == nil {
			db.DB.Create(&models.TmplType{
				Name:     tt.Name,
				Icon:     tt.Icon,
				Category: tt.Category,
			})
		} else if !req.Skip {
			tmplType.Icon = tt.Icon
			tmplType.Name = tt.Name
			db.DB.Save(tmplType)
		}
	}

	for _, tmpl := range data.Tmpls {
		tmplType := models.TmplTypeBy("category", tmpl.Type)
		if tmplType == nil {
			continue
		}
		tmplModel := models.TmplBy("name", tmpl.Name)
		if tmplModel == nil {
			db.DB.Create(&models.Tmpl{
				Name:        tmpl.Name,
				Type:        tmplType.ID,
				Description: tmpl.Description,
				Requirement: tmpl.Requirement,
				Flow:        tmpl.Flow,
				Listed:      tmpl.Listed,
			})
		} else if !req.Skip {
			tmplModel.Name = tmplType.Name
			tmplModel.Type = tmplType.ID
			tmplModel.Description = tmpl.Description
			tmplModel.Requirement = tmpl.Requirement
			tmplModel.Flow = tmpl.Flow
			tmplModel.Listed = tmpl.Listed
			db.DB.Save(tmplModel)
		}
	}

	for _, device := range data.Devices {
		tmpl := models.TmplBy("name", device.Tmpl)
		if tmpl == nil {
			continue
		}
		deviceModel := models.DeviceBy("name", device.Name)
		if deviceModel == nil {
			db.DB.Create(&models.Device{
				Name:   device.Name,
				Tmpl:   tmpl.ID,
				Config: device.Config,
				Node:   device.Node,
			})
		} else if !req.Skip {
			deviceModel.Tmpl = tmpl.ID
			deviceModel.Config = device.Config
			deviceModel.Node = device.Node
			db.DB.Save(deviceModel)
		}
	}

	for _, layout := range data.Layouts {
		layoutModel := models.LayoutBy("name", layout.Name)
		if layoutModel == nil {
			db.DB.Create(&models.Layout{
				Name:      layout.Name,
				Size:      layout.Size,
				Published: layout.Publish,
				Content:   layout.Content,
			})
		} else if !req.Skip {
			layoutModel.Size = layout.Size
			layoutModel.Published = layout.Publish
			layoutModel.Content = layout.Content
			db.DB.Save(layoutModel)
		}
	}

	for _, setting := range data.Settings {
		settingModel := models.QuerySetting(setting.Key)
		if settingModel == nil {
			db.DB.Create(&models.Settings{
				Key:   setting.Key,
				Value: setting.Value,
			})
		} else if !req.Skip {
			settingModel.Value = setting.Value
			db.DB.Save(settingModel)
		}
	}

	c.Bye(gin.H{"data": "ok"})
}
