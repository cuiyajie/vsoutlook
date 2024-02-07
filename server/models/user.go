package models

import (
	"time"

	"github.com/jinzhu/copier"
	"vsoutlook.com/vsoutlook/models/db"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ModelId
	Name     string
	Password string `gorm:"type:varchar"`
	Role     uint8  `gorm:"default:2"`
	ModelTime
}

type UserAsBasic struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role uint8  `json:"role"`
}

func (user User) IsDeleted() bool {
	return false
}

func (user *User) AsBasic() UserAsBasic {
	var result UserAsBasic
	copier.Copy(&result, &user)
	return result
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// this can be slow(~1 seconds), it's by design
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UserBasic(uid string) UserAsBasic {
	user := GetUser(uid)
	return user.AsBasic()
}

func UserList() []UserAsBasic {
	var users []User
	db.DB.Find(&users)
	result := make([]UserAsBasic, 0, len(users))
	for _, u := range users {
		result = append(result, u.AsBasic())
	}
	return result
}

func UsersBasics(uids []string) []UserAsBasic {
	var users []User
	db.DB.Where("id IN ?", uids).Find(&users)
	result := make([]UserAsBasic, 0, len(users))
	for _, u := range users {
		result = append(result, u.AsBasic())
	}
	return result
}

func (user *User) RegisteredDays() int {
	return int(time.Since(user.CreatedAt).Hours() / 24.0)
}
