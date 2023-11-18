package utils

import (
	_ "embed"
	"runtime/debug"
	"strings"
)

//go:embed build-info.txt
var buildInfo string

func BuildInfo() string {
	return buildInfo
}

func BuildInfoVersion() string {
	return strings.Split(buildInfo, " - ")[0]
}

func BuildInfo0() string {
	var (
		commit   = ""
		time     = ""
		modified = false
		result   = ""
	)
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				commit = setting.Value[:10]
			}
			if setting.Key == "vcs.time" {
				time = setting.Value
			}
			if setting.Key == "vcs.modified" {
				modified = setting.Value == "true"
			}
		}
	}
	result = strings.Join([]string{commit, time}, " ")
	if modified {
		result = result + " (modified)"
	}
	return result
}
