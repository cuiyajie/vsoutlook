package config

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

var AppDir = "/app"
var AppDataDir = "/app/data"
var Mode = "app" // "cmd"
var Dev = false
var Testing = false
var DBLogLevel = logger.Warn
var Hostname = "app"
var ClusterID = ""

// 仅用于 log，用于兼容旧 log 格式
var LogSite = ""

var AppSiteDir = "/app/site"

func LoadConfig(filePath string) {

	Hostname, _ = os.Hostname()
	viper.SetConfigType("yaml")
	appConfig, err := os.ReadFile(filePath)
	if err != nil {
		if len(os.Getenv("REDIS_URL")) == 0 {
			panic(err)
		}
		return
	}

	viper.ReadConfig(bytes.NewBuffer(appConfig))
}

func Get(key string) string {
	result := viper.GetString(key)
	if len(result) == 0 {
		return os.Getenv(key)
	}
	return result
}

func Set(key string, val string) {
	viper.Set(key, val)
}

func SetupEnv() {
	ClusterID = os.Getenv("CLUSTER_ID")
	if runtime.GOOS == "darwin" {
		home := os.Getenv("HOME")
		AppDir = home + "/Disposable/vsoutlook-app"
		AppDataDir = AppDir + "/data"
		Dev = true
		DBLogLevel = logger.Info
		ClusterID = "dev"

		// 即使是开发环境，程序不一定总在代码目录下运行，所以配置文件用全局路径
		dir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("err: %s", err)
		}
		fname := filepath.Join(dir, ".config", "vsoutlook-server.yaml")
		if !fileExists(fname) {
			log.Fatalf("config file not found: %s", fname)
		}
		LoadConfig(fname)
		// ─────────────────────────────────────────────────────────────────

		log.SetFlags(log.LstdFlags | log.Llongfile)
	}

	if !Dev {
		gin.SetMode(gin.ReleaseMode)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	fname := filepath.Join(AppDataDir, "config.yaml")

	if fileExists(fname) {
		LoadConfig(fname)
	}

	LogSite = ClusterID
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
