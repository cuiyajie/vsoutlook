package utils

import "github.com/microcosm-cc/bluemonday"

var bluemondayPolicy *bluemonday.Policy

func init() {
	bluemondayPolicy = bluemonday.UGCPolicy()
}

func SafeHTML(input string) string {
	return bluemondayPolicy.Sanitize(input)
}
