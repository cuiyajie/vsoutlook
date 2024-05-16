package models

import (
	"gorm.io/datatypes"
	"vsoutlook.com/vsoutlook/models/db"
)

type Layout struct {
	ModelId
	Name      string                           `gorm:"type:varchar"`
	Size      uint8                            `gorm:"default:1"`
	Published uint8                            `gorm:"default:0"`
	Location  string                           `gorm:"type:varchar"`
	Content   datatypes.JSONSlice[interface{}] `gorm:"type:jsonb"`
	Deleted   uint8                            `gorm:"default:0"`
	ModelTime
}

type LayoutAsBasic struct {
	ID        string                           `json:"id"`
	Name      string                           `json:"name"`
	Size      uint8                            `json:"size"`
	Published uint8                            `json:"published"`
	Location  string                           `json:"location"`
	Content   datatypes.JSONSlice[interface{}] `json:"content,omitempty"`
	CreatedAt int64                            `json:"createdAt"`
	UpdatedAt int64                            `json:"updatedAt"`
}

func (layout Layout) IsDeleted() bool {
	return layout.Deleted > 0
}

func (layout *Layout) AsBasic() LayoutAsBasic {
	var result LayoutAsBasic
	result.ID = layout.ID
	result.Name = layout.Name
	result.Size = layout.Size
	result.Published = layout.Published
	result.Location = layout.Location
	result.Content = layout.Content
	result.CreatedAt = layout.CreatedAt.Unix()
	result.UpdatedAt = layout.UpdatedAt.Unix()
	return result
}

func LayoutList() []LayoutAsBasic {
	var layouts []Layout
	db.DB.Where("deleted=0").Find(&layouts)
	result := make([]LayoutAsBasic, 0, len(layouts))
	for _, t := range layouts {
		result = append(result, t.AsBasic())
	}
	return result
}

func PublishedLayoutList() []Layout {
	var layouts []Layout
	db.DB.Where("published=1 and deleted=0").Find(&layouts)
	return layouts
}
