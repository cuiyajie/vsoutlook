package models

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"vsoutlook.com/vsoutlook/models/db"
)

func Migrate(env string) {
	db.Open(env, "")

	m := gormigrate.New(db.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// 添加新的migration, 参考 https://github.com/go-gormigrate/gormigrate
		{
			ID: "20231118-1457",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					Role uint8 `gorm:"default:1"`
				}
				return tx.AutoMigrate(&User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("users", "role")
			},
		},
		{
			ID: "20231112-1950",
			Migrate: func(tx *gorm.DB) error {
				type Tmpl struct {
					Listed uint8 `gorm:"default:1"`
				}
				return tx.AutoMigrate(&Tmpl{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("tmpls", "listed")
			},
		},
		{
			ID: "20231110-0728",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Tmpl{})
			},
		},
		{
			ID: "20231109-2151",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&TmplType{})
			},
		},
		{
			ID: "20231106-2210",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&db.Idgen{})
			},
		},
	})

	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&db.Idgen{},
			&User{},
			&TmplType{},
			&Tmpl{},
		)
		if err != nil {
			return err
		}
		return nil
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
