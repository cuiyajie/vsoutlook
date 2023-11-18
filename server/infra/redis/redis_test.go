package redis

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"vsoutlook.com/vsoutlook/infra/utils"
)

func TestConnect(t *testing.T) {
	viper.Set("REDIS_URL", "redis://localhost:6379/5")
	ConnectRedis()
	code := utils.SmsCode()
	key := "19900001111"
	SetEx(key, code, time.Duration(30)*time.Minute)
	storedCode := Get(key)
	assert.Equal(t, code, storedCode)
}
