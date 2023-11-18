package db

import (
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"vsoutlook.com/vsoutlook/infra/config"
)

var DB *gorm.DB

func Open(env string, driver string) {
	var err error
	var dbUrl string

	dbUrl = config.Get("DATABASE_URL")
	if env == "test" {
		dbUrl = config.Get("DATABASE_TESTURL")
	}

	var u *url.URL
	if u, err = url.Parse(dbUrl); err != nil {
		log.Printf("invalid db url: %s", dbUrl)
		panic("")
	}

	// k8s 中，DB_PASSWORD 单独管理/传递
	if _, passwordSet := u.User.Password(); !passwordSet {
		if len(os.Getenv("DB_PASSWORD")) > 0 {
			dbUrl = strings.Replace(dbUrl, "@", ":"+os.Getenv("DB_PASSWORD")+"@", 1)
		}
	}

	gormConfig := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	if config.Dev {
		stdout := logger.New(
			// log.New(os.Stdout, "", log.LstdFlags),
			&StrippedLogger{
				log.New(os.Stdout, ">> ", 0),
			},
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  config.DBLogLevel,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		)
		gormConfig.Logger = stdout
	}

	pgConfig := postgres.Open(dbUrl)
	if len(driver) > 0 {
		pgConfig = postgres.New(postgres.Config{
			DSN:        dbUrl,
			DriverName: driver,
		})
	}
	if DB, err = gorm.Open(pgConfig, &gormConfig); err != nil {
		log.Printf("err %#v\n", err)
		panic(err.Error())
	}
	RegisterHooks()
}

func Close() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}

func RegisterHooks() {
	log.Print("db hooks registered")
	DB.Callback().Create().Before("gorm:create").Register("app:set_hash_id_when_create", func(db *gorm.DB) {
		var field = db.Statement.Schema.LookUpField("ID")
		if field != nil {
			if _, isEmpty := field.ValueOf(db.Statement.ReflectValue); isEmpty {
				field.Set(db.Statement.ReflectValue, GenerateID())
			}
		}
	})
}

func JSONField(db *gorm.DB) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
