package models

import (
	"vsoutlook.com/vsoutlook/models/db"
)

type Device struct {
	ModelId
	Name    string `gorm:"type:varchar"`
	Tmpl    string `gorm:"index;type:varchar"`
	Deleted uint8  `gorm:"default:0"`
	ModelTime
}

type DeviceAsBasic struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TmplID       string `json:"tmplID"`
	TmplName     string `json:"tmplName"`
	TmplTypeID   string `json:"tmplTypeID"`
	TmplTypeName string `json:"tmplTypeName"`
	TmplTypeIcon string `json:"tmplTypeIcon"`
}

func (device Device) IsDeleted() bool {
	return false
}

func (device *Device) AsBasic() DeviceAsBasic {
	var result DeviceAsBasic
	result.ID = device.ID
	result.Name = device.Name
	tmpl := ActiveTmpl(device.Tmpl)
	tmplType := ActiveTmplType(tmpl.Type)
	result.TmplID = tmpl.ID
	result.TmplName = tmpl.Name
	result.TmplTypeID = tmplType.ID
	result.TmplTypeName = tmplType.Name
	result.TmplTypeIcon = tmplType.Icon
	return result
}

func DeviceList() []DeviceAsBasic {
	var devices []Device
	db.DB.Where("deleted=0").Find(&devices)
	result := make([]DeviceAsBasic, 0, len(devices))
	for _, t := range devices {
		result = append(result, t.AsBasic())
	}
	return result
}
