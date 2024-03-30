package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"vsoutlook.com/vsoutlook/models/db"
)

type Uint32Slice []uint32
type MapUint32Slice map[string]Uint32Slice
type StringSlice []string
type MapStringSlice map[string]StringSlice

type Node struct {
	ID           string         `gorm:"type:varchar;primary_key"`
	CoreList     string         `gorm:"type:varchar"`
	Allocated    MapUint32Slice `gorm:"type:jsonb"`
	DMAList      string         `gorm:"type:varchar"`
	AllocatedDMA MapStringSlice `gorm:"type:jsonb"`
	VFCount      uint32         `gorm:"type:integer;default:4"`
}

type NodeAsBasic struct {
	ID           string         `json:"id"`
	CoreList     string         `json:"coreList"`
	Allocated    MapUint32Slice `json:"allocated"`
	DMAList      string         `json:"dmaList"`
	AllocatedDMA MapStringSlice `json:"allocatedDMA"`
	VFCount      uint32         `json:"vfCount"`
}

func (node *Node) AsBasic() NodeAsBasic {
	var result NodeAsBasic
	result.ID = node.ID
	result.CoreList = node.CoreList
	result.Allocated = node.Allocated
	result.DMAList = node.DMAList
	result.AllocatedDMA = node.AllocatedDMA
	result.VFCount = node.VFCount
	return result
}

func ActiveNode(id string) *Node {
	var node Node
	if err := db.DB.Where("id = ?", id).First(&node).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("QueryOne err %s\n", err)
		}
		return nil
	}
	return &node
}
