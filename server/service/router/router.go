package router

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"vsoutlook.com/vsoutlook/infra/track"
	"vsoutlook.com/vsoutlook/service/backup"
	"vsoutlook.com/vsoutlook/service/cluster"
	"vsoutlook.com/vsoutlook/service/device"
	"vsoutlook.com/vsoutlook/service/endswt"
	"vsoutlook.com/vsoutlook/service/layout"
	"vsoutlook.com/vsoutlook/service/settings"
	"vsoutlook.com/vsoutlook/service/site"
	"vsoutlook.com/vsoutlook/service/svcinfra"
	"vsoutlook.com/vsoutlook/service/tmpl"
	userSvc "vsoutlook.com/vsoutlook/service/user"
)

func SetupRouter() *gin.Engine {
	track.Setup()
	router := gin.New()
	router.Use(gin.Logger(),
		gin.CustomRecovery(handleRecovery),
		cors.New(corsConfig()),
		logger)
	root := &router.RouterGroup

	var user = new(userSvc.UserService)
	var authenticated = svcinfra.RequireLogin

	{
		guest := group(root, "api")
		post(guest, "login", userSvc.Login)
		get(guest, "internal/init", userSvc.Init)
	}

	{
		guest := group(root, "api")
		post(guest, "logout", user.Logout)
	}

	{
		member := group(root, "api", authenticated)
		post(member, "user/me", user.CurrentUser)
		post(member, "user/list", user.UserList)
		post(member, "user/create", userSvc.CreateUser)
	}

	{
		member := group(root, "api", authenticated)
		// member := group(root, "api")
		post(member, "tmpl_type/list", tmpl.GetTmplTypeList)
		post(member, "tmpl_type/create", tmpl.CreateTmplType)
		post(member, "tmpl/list", tmpl.GetTmplList)
		post(member, "tmpl/info", tmpl.GetTmpl)
		post(member, "tmpl/create", tmpl.CreateTmpl)
		post(member, "tmpl/update", tmpl.UpdateTmpl)
		post(member, "tmpl/delete", tmpl.DeleteTmpl)
		post(member, "tmpl/listed", tmpl.SetTmplListed)
		post(member, "device/list", device.GetDeviceList)
		post(member, "device/detail", device.GetDevice)
		post(member, "device/create", device.CreateDevice)
		post(member, "device/update", device.UpdateDevice)
		post(member, "device/start", device.StartDevice)
		post(member, "device/stop", device.StopDevice)
		post(member, "device/delete", device.DeleteDevice)
	}

	{
		member := group(root, "api", authenticated)
		post(member, "cluster/podPhase", cluster.GetPodsPhase)
		post(member, "cluster/pod.delete", cluster.DeletePod)
		post(member, "cluster/nodes", cluster.GetNodes)
		post(member, "cluster/node.detail", cluster.GetNodeDetail)
		post(member, "cluster/node.update", cluster.UpdateNode)
	}

	{
		member := group(root, "api", authenticated)
		post(member, "settings/list", settings.GetSettings)
		post(member, "settings/update", settings.UpdateSetting)
	}

	{
		member := group(root, "api", authenticated)
		post(member, "layout/list", layout.GetLayoutList)
		post(member, "layout/detail", layout.GetLayout)
		post(member, "layout/create", layout.CreateLayout)
		post(member, "layout/update", layout.UpdateLayout)
		post(member, "layout/delete", layout.DeleteLayout)
		post(member, "layout/publish", layout.PublishLayout)
		post(member, "layout/duplicate", layout.Duplicate)
		post(member, "layout/update.content", layout.UpdateLayoutContent)
	}

	{
		member := group(root, "api", authenticated)
		post(member, "endswt/target", endswt.GetEndSwtTarget)
		post(member, "endswt/switch", endswt.SwitchTarget)
	}

	{
		member := group(root, "api", authenticated)
		post(member, "backup/export", backup.ExportData)
		post(member, "backup/import", backup.ImportData)
	}

	router.NoRoute(h(site.FallbackHandler))
	return router
}

var h = svcinfra.ToHandler
var hs = svcinfra.ToHandlers

func post(group *gin.RouterGroup, path string, handlers ...svcinfra.HandlerFunc) {
	group.POST(path, hs(handlers...)...)
}

func get(group *gin.RouterGroup, path string, handlers ...svcinfra.HandlerFunc) {
	group.GET(path, hs(handlers...)...)
}

func group(parent *gin.RouterGroup, path string, handlers ...svcinfra.HandlerFunc) *gin.RouterGroup {
	return parent.Group(path, hs(handlers...)...)
}

func corsConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	// config.AllowAllOrigins = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	config.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type", "Authorization", "X-RTID", "X-Signature"}
	return config
}

func handleRecovery(c *gin.Context, err any) {
	msg := fmt.Sprintf("panic: %s %s\nrequest: %s\n%s",
		c.Request.RequestURI, err, c.GetString("bodyPreview"), errorStack(4))
	fmt.Println(msg)
}

func errorStack(skip int) []byte {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.HasPrefix(file, "/code/") || strings.HasPrefix(file, "/Users/") {
			fmt.Fprintf(buf, "%s:%d \n", file, line)
		}
	}
	return buf.Bytes()
}
