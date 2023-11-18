package track

import (
	"fmt"
	"strings"

	"github.com/avct/uasurfer"
	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/track/dailylogger"
	"vsoutlook.com/vsoutlook/infra/utils"
	"vsoutlook.com/vsoutlook/models"
)

var daily *dailylogger.Logger

func Setup() {
	if daily == nil {
		daily = dailylogger.NewLogger(config.AppDataDir + "/logs")
	}
}

func Log(info utils.Map) {
	info["server"] = config.Hostname
	info["cluster"] = config.ClusterID
	info[Site] = config.LogSite
	daily.Log(info)
}

func OnUser(obj utils.Map, user *models.User) {
	if user != nil {
		obj[UserID] = user.ID
		obj[RegisteredDays] = user.RegisteredDays()
	} else {
		obj[UserID] = "guest"
	}
}

func OnUserAgent(obj map[string]string, uastr string) {
	ua := uasurfer.Parse(uastr)
	browserName := ua.Browser.Name.StringTrimPrefix()
	if browserName == "IE" && strings.Contains(uastr, " Edg/") {
		browserName = "Edge"
	}

	obj[Browser] = browserName
	obj[BrowserVersion] = ConcatVersion(ua.Browser.Version)
	obj[Platform] = ua.OS.Platform.StringTrimPrefix()
	osName := ua.OS.Name.StringTrimPrefix()
	if osName == "MacOSX" {
		osName = "macOS"
	}
	obj[OperatingSystem] = osName
	obj[OSVersion] = ConcatVersion(ua.OS.Version)
}

// todo: is there a way to make a generic OnUserAgent
// todo: dup code with OnUserAgent, wihout Platform/OS
func OnUserAgent1(obj utils.Map, uastr string) {
	ua := uasurfer.Parse(uastr)
	browserName := ua.Browser.Name.StringTrimPrefix()
	if browserName == "IE" && strings.Contains(uastr, " Edg/") {
		browserName = "Edge"
	}

	obj[Browser] = browserName
	obj[BrowserVersion] = ConcatVersion(ua.Browser.Version)
}

func ConcatVersion(v uasurfer.Version) string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
