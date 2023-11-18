package svcinfra

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/sessions"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models"
)

type HandlerFunc func(*Context)

var Site = "vsoutook"

type Context struct {
	*gin.Context

	Intl      bool
	Lang      int
	UID       string
	RtID      string
	TzOffset  string
	BrowserID string
	Session   *sessions.Session

	User *models.User
}

func NewContext(c *gin.Context) *Context {
	// `${timestamp},${tz},${runtime.browserID},${sig}`,
	sigBundle := strings.Split(c.Request.Header.Get("X-Signature"), ",")
	var tzOffset string
	var bsid string
	var rtid string
	if len(sigBundle) == 4 || len(sigBundle) == 5 {
		tzOffset = sigBundle[1]
		bsid = sigBundle[2]
	}
	if len(sigBundle) == 5 {
		rtid = sigBundle[3]
	}

	ctx := &Context{
		Context:   c,
		TzOffset:  tzOffset,
		BrowserID: bsid,
		RtID:      rtid,
	}

	ctx.Session = GetSession(ctx.Request)
	uid := sessGetString(ctx.Session, SessionUIDKey)
	c.Set(UIDKey, uid)
	ctx.UID = uid
	ctx.User = models.GetUser(ctx.UID)

	return ctx
}

func (c *Context) RenewSession() {
	if cookie, err := c.Request.Cookie(SessionKey); err == nil {
		if cookie.Expires.Sub(time.Now()).Seconds() < 86400*80 {
			sessSetString(c, c.Session, SessionUIDKey, c.UID)
		}
	}
}

func (c *Context) SafeBind(obj any) error {
	return c.ShouldBindBodyWith(obj, binding.JSON)
}

func (c *Context) Bye(obj any) {
	if obj == nil {
		obj = gin.H{}
	}
	c.JSON(http.StatusOK, obj)
	c.Abort()
}

func (c *Context) Save(val any) bool {
	result := models.Save(val)
	if result.Error != nil {
		var id = utils.RandString(16)
		log.Printf("#%s models.Save error: %s", id, result.Error)
		c.GeneralError(fmt.Sprintf("error on saving record #%s", id))
		return false
	}
	return true
}

func (c *Context) CreateModel(val any) bool {
	result := models.Create(val)
	if result.Error != nil {
		var id = utils.RandString(16)
		log.Printf("#%s models.Create error: %s", id, result.Error)
		c.GeneralError(fmt.Sprintf("error on creating record #%s", id))
		return false
	}
	return true
}

func (c *Context) BindOrQuit(obj any) bool {
	if err := c.SafeBind(obj); err != nil {
		c.InvalidParams("")
		return false
	}
	return true
}

func ToHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ictx, ok := c.Get(CtxServiceContextKey); ok {
			handler(ictx.(*Context))
		} else {
			ictx := NewContext(c)
			c.Set(CtxServiceContextKey, ictx)
			handler(ictx)
		}
	}
}

func ToHandlers(handlers ...HandlerFunc) (result []gin.HandlerFunc) {
	for _, f := range handlers {
		result = append(result, ToHandler(f))
	}
	return
}

const (
	CtxBaseInputKey      = "base"
	CtxServiceContextKey = "service-context"
	CtxReqParams         = "req-params"
)

// 内部记录出错信息，tests 用
var QuitPoint = ""
var Testing = false

const UIDKey = "UID"
