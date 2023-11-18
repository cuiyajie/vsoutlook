package user

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/service/svcinfra"
)

type UserService struct{}

type UserReq struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const intlInvalidLogin = "请输入正确的用户名和密码。"
const intlUnknownError = "Unknown error."

func CreateUser(c *svcinfra.Context) {
	var req UserReq
	c.ShouldBindJSON(&req)
	log.Printf("req %#v\n", req)

	if len(req.Password)*len(req.Username) == 0 {
		c.GeneralError("Invalid Input.")
		return
	}

	req.Username = strings.TrimSpace(req.Username)

	user := models.UserBy("name", req.Username)
	if user != nil {
		c.GeneralError("Username already exits.")
		return
	}

	if len(req.Password) < 3 {
		c.GeneralError("Password too short.")
		return
	}

	hashedPwd, err := models.HashPassword(req.Password)
	if err != nil {
		c.GeneralError(intlUnknownError)
		return
	}

	var newUser = models.User{
		Name:     req.Username,
		Password: hashedPwd,
	}

	result := models.Create(&newUser)
	if result.Error != nil {
		fmt.Errorf("create user failed: %w", result.Error)
		c.GeneralError(intlUnknownError)
		return
	}
	user = &newUser

	c.Bye(gin.H{"result": "ok"})
}

func Login(c *svcinfra.Context) {
	var req UserReq
	c.ShouldBindJSON(&req)
	req.Username = strings.TrimSpace(req.Username)
	if len(req.Username) == 0 || len(req.Password) == 0 {
		c.GeneralError(intlInvalidLogin)
		return
	}
	user := models.UserBy("name", req.Username)
	if user == nil {
		c.GeneralError(intlInvalidLogin)
		return
	}

	if !models.CheckPasswordHash(req.Password, user.Password) {
		c.GeneralError(intlInvalidLogin)
		return
	}

	if !svcinfra.SetSessionUser(c, user.ID) {
		c.GeneralError(intlInvalidLogin)
		return
	}
	c.Bye(gin.H{"user": user.AsBasic()})
}

func (us *UserService) CurrentUser(c *svcinfra.Context) {
	user := c.User
	c.Bye(gin.H{"user": user.AsBasic()})
}

func (us *UserService) Logout(c *svcinfra.Context) {
	// sess := getSession(c)
	// sess.Options.MaxAge = -1 // 设置 MaxAge=-1 对大部分 browser 有用，对微信内置浏览器不管用
	// 微信内置浏览器会继续用之前服务器设置过的 cookie
	svcinfra.SetSessionUser(c, "guest")
	c.Bye(nil)
}

func (us *UserService) UserList(c *svcinfra.Context) {
	c.Bye(gin.H{"users": models.UserList()})
}
