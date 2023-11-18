package redis

import "time"

// const min30 = 1800 * time.Second

func Notified_TimeWindowKey(t time.Time) string {
	n := t.Unix()
	start := n - n%1800
	return "notified-" + time.Unix(start, 0).Format("20060102-1504")
}

func Notified_CurrentTimeWindowKey() string {
	return Notified_TimeWindowKey(time.Now())
}

func FilesToDeleteSources() string {
	return "files-to-be-cleared"
}
