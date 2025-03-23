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
			ID: "20250315-1742",
			Migrate: func(tx *gorm.DB) error {
				// 为 tmpl_type增加 deleted 字段
				type TmplType struct {
					AppCategory string `gorm:"type:varchar"`
				}
				return tx.AutoMigrate(&TmplType{})
			},
			Rollback: func(tx *gorm.DB) error {
				// drop app category column
				return tx.Migrator().DropColumn("tmpl_types", "app_category")
			},
		},
		{
			ID: "20250315-1054",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Nic{})
			},
		},
		{
			ID: "20250311-2324",
			Migrate: func(tx *gorm.DB) error {
				// 为 tmpl_type增加 deleted 字段
				type Node struct {
					CoreList     string         `gorm:"type:varchar"`
					Allocated    MapUint32Slice `gorm:"type:jsonb"`
					DMAList      string         `gorm:"type:varchar"`
					AllocatedDMA MapStringSlice `gorm:"type:jsonb"`
					VFCount      uint32         `gorm:"type:integer;default:4"`
				}
				if err := tx.Migrator().DropColumn(&Node{}, "core_list"); err != nil {
					log.Printf("DropColumn core_list err %s\n", err)
					return err
				}
				if err := tx.Migrator().DropColumn(&Node{}, "dma_list"); err != nil {
					log.Printf("DropColumn dma_list err %s\n", err)
					return err
				}
				if err := tx.Migrator().DropColumn(&Node{}, "vf_count"); err != nil {
					log.Printf("DropColumn vf_count err %s\n", err)
					return err
				}
				if err := tx.Migrator().DropColumn(&Node{}, "allocated"); err != nil {
					log.Printf("DropColumn allocated err %s\n", err)
					return err
				}
				if err := tx.Migrator().DropColumn(&Node{}, "allocated_dma"); err != nil {
					log.Printf("DropColumn allocated_dma err %s\n", err)
					return err
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				type Node struct {
					CoreList     string         `gorm:"type:varchar"`
					DMAList      string         `gorm:"type:varchar"`
					AllocatedDMA MapStringSlice `gorm:"type:jsonb"`
					VFCount      uint32         `gorm:"type:integer;default:4"`
				}
				if err := tx.AutoMigrate(&Node{}); err != nil {
					log.Printf("AutoMigrate Node err %s\n", err)
					return err
				}
				return nil
			},
		},
		{
			ID: "20250310-2053",
			Migrate: func(tx *gorm.DB) error {
				// 为 tmpl_type增加 deleted 字段
				type TmplType struct {
					Deleted uint8 `gorm:"default:0"`
				}
				return tx.AutoMigrate(&TmplType{})
			},
			Rollback: func(tx *gorm.DB) error {
				// drop app category column
				return tx.Migrator().DropColumn("tmpl_types", "deleted")
			},
		},
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
	tmplTypes := []TmplType{
		{Name: "编解码", Icon: "codec.svg", Category: "codec", AppCategory: "codec", Deleted: 1},
		{Name: "上下变换", Icon: "udx.svg", Category: "udx", AppCategory: "udx", Deleted: 1},
		{Name: "多画面", Icon: "mv.svg", Category: "multiv", AppCategory: "mv", Deleted: 1},
		{Name: "切换台", Icon: "switch.svg", Category: "swt", AppCategory: "swt", Deleted: 1},
		{Name: "末级切换", Icon: "endswitch.svg", Category: "endswt", AppCategory: "swt", Deleted: 1},
		{Name: "录放机", Icon: "recorder.svg", Category: "recorder", AppCategory: "rc"},
		{Name: "全媒体网关", Icon: "media_gateway.svg", Category: "media_gateway", AppCategory: "mg"},
		{Name: "多画面", Icon: "mv.svg", Category: "mv", AppCategory: "mv"},
		{Name: "播出切换台", Icon: "bcswitch.svg", Category: "bcswt", AppCategory: "swt"},
		{Name: "制作切换台", Icon: "switch_v2.svg", Category: "makeswt", AppCategory: "swt"},
		{Name: "新媒体切换台", Icon: "new_media.svg", Category: "nmswt", AppCategory: "swt"},
	}
	for _, tmplType := range tmplTypes {
		// 检查是否已存在, 以 Category 为主键, 使用 db_helpers.go 里面的 QueryOne
		tmplTypeRecord := QueryOne[TmplType]("category", tmplType.Category)
		if tmplTypeRecord == nil && tmplType.Deleted != 1 {
			// 不存在则创建
			Create(&tmplType)
			log.Printf("Create tmplType: %s", tmplType.Category)
		} else {
			// 存在则更新
			tmplTypeRecord.Name = tmplType.Name
			tmplTypeRecord.Icon = tmplType.Icon
			tmplTypeRecord.AppCategory = tmplType.AppCategory
			tmplTypeRecord.Deleted = tmplType.Deleted
			Save(&tmplTypeRecord)
			log.Printf("Update tmplType: %s", tmplType.Category)
		}
	}
}
