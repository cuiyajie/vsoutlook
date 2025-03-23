package models

import (
	"vsoutlook.com/vsoutlook/models/db"
)

type Device struct {
	ModelId
	Name    string `gorm:"type:varchar"`
	Tmpl    string `gorm:"index;type:varchar"`
	AppName string `gorm:"type:varchar"`
	Config  string `gorm:"type:varchar"`
	Node    string `gorm:"type:varchar"`
	Deleted uint8  `gorm:"default:0"`
	SeedID  string `gorm:"type:varchar"`
	ModelTime
}

type DeviceAsBasic struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	AppName          string `json:"appName"`
	TmplID           string `json:"tmplID"`
	TmplName         string `json:"tmplName"`
	TmplTypeID       string `json:"tmplTypeID"`
	TmplTypeName     string `json:"tmplTypeName"`
	TmplTypeIcon     string `json:"tmplTypeIcon"`
	TmplTypeCategory string `json:"tmplTypeCategory"`
	Node             string `json:"node"`
	CreatedAt        int64  `json:"createdAt"`
	UpdatedAt        int64  `json:"updatedAt"`
}

type DeviceAsDetail struct {
	DeviceAsBasic
	Config string `json:"config"`
}

func (device Device) IsDeleted() bool {
	return device.Deleted > 0
}

func (device *Device) AsBasic() DeviceAsBasic {
	var result DeviceAsBasic
	result.ID = device.ID
	result.Name = device.Name
	tmpl := ActiveTmpl(device.Tmpl)
	if tmpl != nil {
		tmplType := ActiveTmplType(tmpl.Type)
		if tmplType != nil {
			result.TmplID = tmpl.ID
			result.TmplName = tmpl.Name
			result.TmplTypeID = tmplType.ID
			result.TmplTypeName = tmplType.Name
			result.TmplTypeCategory = tmplType.Category
			result.TmplTypeIcon = tmplType.Icon
		}
	}
	result.CreatedAt = device.CreatedAt.Unix()
	result.UpdatedAt = device.UpdatedAt.Unix()
	result.AppName = device.AppName
	result.Node = device.Node
	return result
}

func (device *Device) AsDetail() DeviceAsDetail {
	result := device.AsBasic()
	var detail DeviceAsDetail
	detail.DeviceAsBasic = result
	detail.Config = device.Config
	return detail
}

func DeviceList() []DeviceAsBasic {
	var devices []Device
	db.DB.Where("deleted=0").Find(&devices)
	result := []DeviceAsBasic{}
	for _, t := range devices {
		basic := t.AsBasic()
		if basic.TmplID != "" {
			result = append(result, basic)
		}
	}
	return result
}

func GetDevicesByNode(node string) []DeviceAsBasic {
	var devices []Device
	db.DB.Where("deleted=0 AND node=?", node).Find(&devices)
	result := make([]DeviceAsBasic, 0, len(devices))
	for _, t := range devices {
		basic := t.AsBasic()
		if basic.TmplID != "" {
			result = append(result, basic)
		}
	}
	return result
}
