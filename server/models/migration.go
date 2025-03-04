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
			ID: "20240411-2200",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Layout{})
			},
		},
		{
			ID: "20240330-1715",
			Migrate: func(tx *gorm.DB) error {
				type Node struct {
					AllocatedDMA MapStringSlice `gorm:"type:jsonb"`
				}
				return tx.AutoMigrate(&Node{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("nodes", "allocated_dma")
			},
		},
		{
			ID: "20240314-2149",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Settings{})
			},
		},
		{
			ID: "20240311-2059",
			Migrate: func(tx *gorm.DB) error {
				type Device struct {
					SeedID string `gorm:"type:varchar"`
				}
				return tx.AutoMigrate(&Device{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("devices", "seed_id")
			},
		},
		{
			ID: "20240304-2143",
			Migrate: func(tx *gorm.DB) error {
				type Node struct {
					VFCount uint32 `gorm:"type:integer;default:4"`
				}
				return tx.AutoMigrate(&Node{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("nodes", "vf_count")
			},
		},
		{
			ID: "20240302-1329",
			Migrate: func(tx *gorm.DB) error {
				type Device struct {
					Node string `gorm:"type:varchar"`
				}
				return tx.AutoMigrate(&Device{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("devices", "node")
			},
		},
		{
			ID: "20240229-2146",
			Migrate: func(tx *gorm.DB) error {
				type Node struct {
					DMAList string `gorm:"type:varchar"`
				}
				return tx.AutoMigrate(&Node{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("nodes", "dma_list")
			},
		},
		{
			ID: "20240225-1048",
			Migrate: func(tx *gorm.DB) error {
				type Device struct {
					Config string `gorm:"type:varchar"`
				}
				return tx.AutoMigrate(&Device{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("devices", "config")
			},
		},
		{
			ID: "20240222-2253",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Node{})
			},
		},
		{
			ID: "20240204-2221",
			Migrate: func(tx *gorm.DB) error {
				type Device struct {
					AppName string `gorm:"type:varchar"`
				}
				return tx.AutoMigrate(&Device{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("devices", "app_name")
			},
		},
		{
			ID: "20240204-2136",
			Migrate: func(tx *gorm.DB) error {
				type TmplType struct {
					Category string `gorm:"type:varchar"`
				}
				return tx.AutoMigrate(&TmplType{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("tmpl_types", "category")
			},
		},
		{
			ID: "20231125-1851",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Device{})
			},
		},
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
			&Device{},
			&Node{},
			&Settings{},
			&Layout{},
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

func SetInitialData() {
	var count int64
	db.DB.Model(&TmplType{}).Count(&count)
	if count == 0 {
		tmplTypes := []TmplType{
			{Name: "编解码", Icon: "codec.svg", Category: "codec"},
			{Name: "上下变换", Icon: "udx.svg", Category: "udx"},
			{Name: "多画面", Icon: "mv.svg", Category: "multiv"},
			{Name: "切换台", Icon: "switch.svg", Category: "swt"},
			{Name: "播出切换台", Icon: "bcswitch.svg", Category: "bcswt"},
			{Name: "末级切换", Icon: "endswitch.svg", Category: "endswt"},
		}
		for _, tmplType := range tmplTypes {
			db.DB.Create(&tmplType)
		}
		log.Print("Initial data for tmpl_types created")
	} else {
		log.Print("Initial data for tmpl_types already exists")
	}
}
