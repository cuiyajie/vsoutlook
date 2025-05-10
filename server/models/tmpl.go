package models

import (
	"github.com/jinzhu/copier"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models/db"
)

type NicConfig struct {
	DPDKCpu int `json:"dpdkCpu"`
	SDKCpu  int `json:"sdkCpu"`
	DMA     int `json:"dma"`
}

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
	Cpu                  string      `json:"cpu"`
	CpuNum               int         `json:"cpuNum"`
	CpuCore              string      `json:"cpuCore"`
	HugePage             int         `json:"hugePage"`
	Memory               int         `json:"memory"`
	Disk                 string      `json:"disk"`
	Gpu                  string      `json:"gpu"`
	InputBand            string      `json:"inputBand"`
	OutputBand           string      `json:"outputBand"`
	Network              string      `json:"network"`
	Chart                string      `json:"chart"`
	LogLevel             uint8       `json:"logLevel"`
	RepairRecvFrame      bool        `json:"repairRecvFrame"`
	RepairSendFrame      bool        `json:"repairSendFrame"`
	HostNetwork          bool        `json:"hostNetwork"`
	UtfOffset            uint8       `json:"utfOffset"`
	RecvAVFrameNodeCount int         `json:"recvAVFrameNodeCount"`
	SendAVFrameNodeCount int         `json:"sendAVFrameNodeCount"`
	RecvframeCount       int         `json:"recvFrameCount"`
	ImageTag             string      `json:"imageTag"`
	MaxRateMbpsByCore    int         `json:"maxRateMbpsByCore"`
	ReceiveSessions      int         `json:"receiveSessions"`
	NicCount             int         `json:"nicCount"`
	NicConfig            []NicConfig `json:"nicConfig"`
	Shm                  int         `json:"shm"`
	CoreType             string      `json:"coreType"`
	TxShare              int         `json:"txShare"`
	RxShare              int         `json:"rxShare"`
}

// coreType: txrx, dpdk, none

type TmplAsBasic struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	TypeName     string `json:"typeName"`
	TypeCategory string `json:"typeCategory"`
	Description  string `json:"description"`
	Listed       int    `json:"listed"`
}

func (tmpl Tmpl) IsDeleted() bool {
	return tmpl.Deleted > 0
}

func (tmpl *Tmpl) AsBasic() TmplAsBasic {
	var result TmplAsBasic
	tmplType := ActiveTmplType(tmpl.Type)
	if tmplType == nil {
		return result
	}
	copier.Copy(&result, &tmpl)
	result.TypeName = tmplType.Name
	result.TypeCategory = tmplType.Category
	return result
}

func (tmpl *Tmpl) AsDetail() any {
	tmplType := ActiveTmplType(tmpl.Type)
	var result utils.Map
	if tmplType == nil {
		return result
	}
	result = utils.Map{
		"id":           tmpl.ID,
		"name":         tmpl.Name,
		"type":         tmpl.Type,
		"typeName":     tmplType.Name,
		"typeCategory": tmplType.Category,
		"description":  tmpl.Description,
		"listed":       tmpl.Listed,
	}
	if tmpl.Requirement.NicConfig == nil {
		tmpl.Requirement.NicConfig = []NicConfig{}
	}
	if tmpl.Requirement.CoreType == "" {
		tmpl.Requirement.CoreType = "txrx"
	}
	result["requirement"] = tmpl.Requirement
	result["flow"] = tmpl.Flow
	return result
}

func TmplList() []TmplAsBasic {
	var tmpls []Tmpl
	db.DB.Where("deleted=0").Find(&tmpls)
	result := []TmplAsBasic{}
	for _, t := range tmpls {
		nicb := t.AsBasic()
		if nicb.ID != "" {
			result = append(result, nicb)
		}
	}
	return result
}
