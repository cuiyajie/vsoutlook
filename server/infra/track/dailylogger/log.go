package dailylogger

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"vsoutlook.com/vsoutlook/infra/config"
	"vsoutlook.com/vsoutlook/infra/utils"
)

type Logger struct {
	dir        string
	lastSwitch time.Time
	f          *os.File
	actual     *log.Logger
}

func NewLogger(dir string) *Logger {
	if len(dir) < 2 {
		log.Fatalf("Invalid log dir: %s", dir)
	}
	if !config.Testing {
		utils.EnsureDirExists(dir)
	}
	l := &Logger{dir: dir}
	l.Switch()
	return l
}

func (l *Logger) Log(info utils.Map) {
	now := time.Now()
	delete(info, "token") // not sure where is "token" added, but it appears in some logs
	since := now.Hour()*3600 + now.Minute()*60 + now.Second()
	if int64(now.Sub(l.lastSwitch)/time.Second) > int64(since) {
		l.Switch()
	}

	var hasTime = false
	if v, exist := info["at"]; exist {
		var v1 int64 = 0
		if vi64, ok := v.(int64); ok {
			v1 = vi64
		} else if vf64, ok := v.(float64); ok {
			// float64 is from JSON parse
			v1 = int64(vf64)
		}

		t := time.Unix(v1, 0)
		if time.Since(t).Seconds() < 86400*180 {
			info["at"] = t.Format(time.RFC3339)
			hasTime = true
		}
	}

	if !hasTime {
		info["at"] = now.Format(time.RFC3339)
	}
	result, err := json.Marshal(&info)
	if err != nil {
		log.Printf("log err: %s", err)
		return
	}
	if l.f == nil || l.actual == nil {
		log.Printf("%s", result)
		return
	}
	l.actual.Printf("%s", result)
}

func (l *Logger) Switch() {
	if l.f != nil {
		l.f.Close()
	}

	now := time.Now()
	logFileName := now.Format("2006-01-02") + ".log"
	targetFile := filepath.Join(l.dir, logFileName)
	log.Printf("will log to %s", targetFile)

	f, err := os.OpenFile(targetFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	l.actual = log.New(f, "", 0)
	l.f = f
	l.lastSwitch = now
}
