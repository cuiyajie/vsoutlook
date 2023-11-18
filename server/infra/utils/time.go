package utils

import (
	"time"
)

func Now() int64 {
	return time.Now().Unix()
}

func NowMsec() int64 {
	return time.Now().UnixMilli()
}
