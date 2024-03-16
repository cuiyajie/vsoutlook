package models

import (
	"errors"
	"log"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"vsoutlook.com/vsoutlook/models/db"
)

type Settings struct {
	ModelId
	Key   string `gorm:"type:varchar"`
	Value string `gorm:"type:varchar"`
}

type SettingsAsBasic struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (settings *Settings) AsBasic() SettingsAsBasic {
	var result SettingsAsBasic
	copier.Copy(&result, &settings)
	return result
}

func GetSettings() map[string]string {
	var settings []Settings
	db.DB.Find(&settings)
	result := make(map[string]string, 0)
	for _, st := range settings {
		result[st.Key] = st.Value
	}
	return result
}

func QuerySetting(key string) *Settings {
	var t Settings
	if err := db.DB.Where("key = ?", key).First(&t).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("QueryOne err %s\n", err)
		}
		return nil
	}
	return &t
}
