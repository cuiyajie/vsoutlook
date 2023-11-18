package models

import (
	"time"
)

type ModelTime struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

type ModelTimeIndex struct {
	CreatedAt time.Time `json:"created_at" gorm:"index;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}

type ModelId struct {
	AutoID uint   `gorm:"primary_key;auto_increment;not null"`
	ID     string `json:"id" gorm:"uniqueIndex;type:varchar;not null"`
}

type Model interface {
	TableName() string
	Deleteable
}

type Deleteable interface {
	IsDeleted() bool
	// MarkDeleted() bool
}

type Basicable interface {
	AsBasic() any
}

type BasicableFor interface {
	AsBasicFor(uid string) any
}

type HasUserIDs interface {
	UserIDs() []string
}

func filterOutDeleted[T Deleteable](rows []T) []T {
	n := 0
	for _, row := range rows {
		if !row.IsDeleted() {
			rows[n] = row
			n++
		}
	}
	return rows[:n]
}
