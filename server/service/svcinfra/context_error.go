package svcinfra

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrUnauthorized  = "unauthorized"
	ErrInvalidParams = "invalid_params"
	ErrGeneral       = "error"
	ErrRequireLogin  = "require_login"
)

var fallbackMessages = map[string]string{
	ErrUnauthorized:  "您没有权限进行此项操作，请检查。",
	ErrInvalidParams: "参数不正确。",
	ErrGeneral:       "遇到错误，可能是您没有权限进行此项操作，请检查。",
	ErrRequireLogin:  "请登录后再试",
}

func (c *Context) InvalidParams(msg string) {
	c.Error(ErrInvalidParams, msg)
}

func (c *Context) NoPermission(msg string) {
	c.Error(ErrUnauthorized, msg)
}

func (c *Context) RequireLogin(msg string) {
	c.Error(ErrRequireLogin, msg)
}

func (c *Context) GeneralError(msg string) {
	c.Error(ErrGeneral, msg)
}

func (c *Context) ErrorWithID(id string, msg string) {
	c.Error(ErrGeneral, fmt.Sprintf("#%s %s", msg, id))
}

func (c *Context) Error(name string, msg string) {
	if msg == "" {
		msg = fallbackMessages[name]
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   name,
		"message": msg,
	})
	c.Abort()
}
