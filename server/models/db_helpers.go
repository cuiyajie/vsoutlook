package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/models/db"
)

func Save(value any) (tx *gorm.DB) {
	return db.DB.Save(value)
}

func Create(value any) (tx *gorm.DB) {
	return db.DB.Create(value)
}

func Updates(value any) (tx *gorm.DB) {
	return db.DB.Updates(value)
}

func ReadyDB() {
	db.Open("", "")
}

func ReadyMaintainTestDB() {
	db.Open("test", "")
}

func ReadyTestDB() {
	config.DBLogLevel = logger.Info
	txdb.Register("txdb", "postgres", config.Get("DATABASE_TESTURL"))
	db.Open("test", "txdb")
}

func Get[T Model](id string) *T {
	if len(id) == 0 {
		return nil
	}
	return QueryOne[T]("id", id)
}

func Active[T Model](id string) *T {
	t := Get[T](id)
	if t == nil || (*t).IsDeleted() {
		return nil
	}
	return t
}

func QueryOne[T Model](column string, value any) *T {
	var t T
	qs := fmt.Sprintf("%s = ?", column)
	if err := db.DB.Where(qs, value).First(&t).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("QueryOne err %s\n", err)
		}
		return nil
	}
	return &t
}

func QueryMulti[T Model](column string, value any) []*T {
	var t []*T
	qs := fmt.Sprintf("%s = ?", column)
	if err := db.DB.Where(qs, value).Find(&t).Error; err != nil {
		log.Printf("QueryMulti err %s\n", err)
		t = t[:0]
	}
	return t
}

// QueryOne4[User]("email", "abcd@gmail.com", "mobile", "18600")
func QueryOne4[T Model](conditions ...any) *T {
	l := len(conditions)
	if l < 2 || l%2 != 0 {
		panic("Invalid parameters for QueryOne")
	}

	tx := db.DB
	i := 0
	for i < l {
		tx = tx.Where(fmt.Sprintf("%s = ?", conditions[i].(string)), conditions[i+1])
		i = i + 2
	}

	var t T
	if err := tx.First(&t).Error; err != nil {
		log.Printf("QueryOne err %s\n", err)
		return nil
	}
	return &t
}

func ByIDs[T Model](ids []string) []*T {
	var t []*T
	if len(ids) == 0 {
		return t
	}
	if err := db.DB.Where("id IN ?", ids).Find(&t).Error; err != nil {
		log.Printf("ByIDs err %s\n", err)
	}
	return t
}

func QueryMulti2[T Model](t *[]T, column string, value any) bool {
	qs := fmt.Sprintf("%s = ?", column)
	if err := db.DB.Where(qs, value).Find(&t).Error; err != nil {
		log.Printf("QueryMulti err %s\n", err)
		return false
	}
	return true
}

func QueryMulti3_0[T Model](order, column string, value any) []*T {
	var t []*T
	qs := fmt.Sprintf("%s = ?", column)
	if err := db.DB.Where(qs, value).Order(order).Find(&t).Error; err != nil {
		log.Printf("QueryMulti3 err %s\n", err)
		t = t[:0]
	}
	return t
}

func QueryMulti3[T Model](order string, conditions ...any) []*T {
	l := len(conditions)
	if l < 2 || l%2 != 0 {
		panic("Invalid parameters for QueryMulti4")
	}

	tx := db.DB
	i := 0
	for i < l {
		tx = tx.Where(fmt.Sprintf("%s = ?", conditions[i].(string)), conditions[i+1])
		i = i + 2
	}

	var t []*T
	if err := tx.Order(order).Find(&t).Error; err != nil {
		log.Printf("QueryMulti3 err %s\n", err)
		return nil
	}
	return t
}

func QueryMulti4[T Model](conditions ...any) []*T {
	l := len(conditions)
	if l < 2 || l%2 != 0 {
		panic("Invalid parameters for QueryMulti4")
	}

	tx := db.DB
	i := 0
	for i < l {
		tx = tx.Where(fmt.Sprintf("%s = ?", conditions[i].(string)), conditions[i+1])
		i = i + 2
	}

	var t []*T
	if err := tx.Find(&t).Error; err != nil {
		log.Printf("QueryMulti4 err %s\n", err)
		return nil
	}
	return t
}
