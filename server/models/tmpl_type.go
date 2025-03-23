package models

import (
	"github.com/jinzhu/copier"
	"vsoutlook.com/vsoutlook/models/db"
)

type TmplType struct {
	ModelId
	Name        string `gorm:"type:varchar"`
	Icon        string `gorm:"type:varchar"`
	Category    string `gorm:"type:varchar"`
	AppCategory string `gorm:"type:varchar"`
	Deleted     uint8  `gorm:"default:0"`
}

type TmplTypeAsBasic struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Category    string `json:"category"`
	AppCategory string `json:"appCategory"`
}

func (tmplType *TmplType) AsBasic() TmplTypeAsBasic {
	var result TmplTypeAsBasic
	copier.Copy(&result, &tmplType)
	return result
}

func (tmplType TmplType) IsDeleted() bool {
	return tmplType.Deleted > 0
}

func TmplTypeList() []TmplTypeAsBasic {
	var tmplTypes []TmplType
	db.DB.Where("deleted=0").Find(&tmplTypes)
	result := make([]TmplTypeAsBasic, 0, len(tmplTypes))
	for _, tt := range tmplTypes {
		result = append(result, tt.AsBasic())
	}
	return result
}
