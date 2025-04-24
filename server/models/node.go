package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"vsoutlook.com/vsoutlook/models/db"
)

type Node struct {
	ID string `gorm:"type:varchar;primary_key"`
}

type NodeAsBasic struct {
	ID string `json:"id"`
}

func (node *Node) AsBasic() NodeAsBasic {
	var result NodeAsBasic
	result.ID = node.ID
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

func EnsureNode(id string) *Node {
	var node Node
	if err := db.DB.Where("id = ?", id).First(&node).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("QueryOne err %s\n", err)
		}
		node.ID = id
		db.DB.Save(&node)
	}
	return &node
}

// get nics by node id
func (node *Node) Nics() []Nic {
	var nics []Nic
	db.DB.Model(&Nic{}).Where("node_id=? and deleted=0", node.ID).Order("position asc").Find(&nics)
	return nics
}
