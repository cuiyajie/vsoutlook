package db

import (
	"log"
	"time"

	"github.com/rs/xid"
	"vsoutlook.com/vsoutlook/infra/utils"
)

type Idgen struct {
	ID string `json:"id" gorm:"type:varchar"`
	At int64  `json:"at" gorm:"type: bigint"`
}

func (Idgen) TableName() string {
	return "idgen"
}

func GenerateID() string {
	var i int
	for {
		candidate := utils.RandString(8)
		now := time.Now().Unix()
		result := DB.Exec(`INSERT into idgen (id, at) VALUES (?, ?)`, candidate, now)
		if result.Error != nil {
			log.Printf("GenerateID insert err: %s", result.Error)
			if i > 5 {
				return xid.New().String()
			}

			i = i + 1
			time.Sleep(time.Second)
			continue
		}
		return candidate
	}
}

// return true if GenerateID works OK
// in case rand not seeded
func CheckGenerateID() bool {
	candidate := utils.RandString(8)
	var result struct {
		ID string `gorm:"column:id"`
	}
	DB.Raw("SELECT id FROM idgen WHERE id = ?", candidate).Scan(&result)
	if result.ID == candidate {
		log.Printf("GenerateID got duplicate id: %s", candidate)
		return false
	}
	return true
}
