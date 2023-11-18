package models

import (
	"github.com/jinzhu/copier"
	"vsoutlook.com/vsoutlook/models/db"
)

type TmplType struct {
	ModelId
	Name string
	Icon string
}

type TmplTypeAsBasic struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (tmplType *TmplType) AsBasic() TmplTypeAsBasic {
	var result TmplTypeAsBasic
	copier.Copy(&result, &tmplType)
	return result
}

func (tmplType TmplType) IsDeleted() bool {
	return false
}

func TmplTypeList() []TmplTypeAsBasic {
	var tmplTypes []TmplType
	db.DB.Find(&tmplTypes)
	result := make([]TmplTypeAsBasic, 0, len(tmplTypes))
	for _, tt := range tmplTypes {
		result = append(result, tt.AsBasic())
	}
	return result
}
