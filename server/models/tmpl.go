package models

import (
	"github.com/jinzhu/copier"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models/db"
)

type Tmpl struct {
	ModelId
	Name        string          `gorm:"type:varchar"`
	Type        string          `gorm:"index;type:varchar"`
	Requirement TmplRequirement `gorm:"column:requirement;type:jsonb"`
	Flow        string
	Description string
	Listed      uint8 `gorm:"default:1"`
	Deleted     uint8 `gorm:"default:0"`
	ModelTime
}

type TmplRequirement struct {
	Cpu                  string `json:"cpu,omitempty"`
	CpuNum               string `json:"cpuNum,omitempty"`
	CpuCore              string `json:"cpuCore,omitempty"`
	HugePage             string `json:"hugePage,omitempty"`
	Memory               string `json:"memory,omitempty"`
	Disk                 string `json:"disk,omitempty"`
	Gpu                  string `json:"gpu,omitempty"`
	InputBand            string `json:"inputBand,omitempty"`
	OutputBand           string `json:"outputBand,omitempty"`
	Network              string `json:"network,omitempty"`
	Chart                string `json:"chart,omitempty"`
	LogLevel             uint8  `json:"logLevel"`
	RepairRecvFrame      bool   `json:"repairRecvFrame"`
	RepairSendFrame      bool   `json:"repairSendFrame"`
	DMAList              string `json:"dmaList,omitempty"`
	HostNetwork          bool   `json:"hostNetwork"`
	UtfOffset            uint8  `json:"utfOffset"`
	RecvAVErameNodeCount int    `json:"recvAVErameNodeCount"`
	SendAVErameNodeCount int    `json:"sendAVErameNodeCount"`
	RecvframeCnt         int    `json:"recvframeCnt"`
	MaxRateMbpsByCore    int    `json:"maxRateMbpsByCore"`
}

type TmplAsBasic struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	TypeName    string `json:"typeName"`
	Description string `json:"description"`
	Listed      int    `json:"listed"`
}

func (tmpl Tmpl) IsDeleted() bool {
	return false
}

func (tmpl *Tmpl) AsBasic() TmplAsBasic {
	var result TmplAsBasic
	copier.Copy(&result, &tmpl)
	tmplType := ActiveTmplType(tmpl.Type)
	result.TypeName = tmplType.Name
	return result
}

func (tmpl *Tmpl) AsDetail() any {
	tmplType := ActiveTmplType(tmpl.Type)
	result := utils.Map{
		"id":          tmpl.ID,
		"name":        tmpl.Name,
		"type":        tmpl.Type,
		"typeName":    tmplType.Name,
		"description": tmpl.Description,
	}
	result["requirement"] = tmpl.Requirement
	result["flow"] = tmpl.Flow
	return result
}

func TmplList() []TmplAsBasic {
	var tmpls []Tmpl
	db.DB.Where("deleted=0").Find(&tmpls)
	result := make([]TmplAsBasic, 0, len(tmpls))
	for _, t := range tmpls {
		result = append(result, t.AsBasic())
	}
	return result
}
