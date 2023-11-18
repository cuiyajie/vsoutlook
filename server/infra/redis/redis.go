package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"vsoutlook.com/vsoutlook/infra/config"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func newRDB() *redis.Client {
	u, err := url.Parse(config.Get("REDIS_URL"))
	if err != nil {
		log.Printf(fmt.Errorf("parse redis url: %w", err).Error())
		panic("")
	}

	password, _ := u.User.Password()
	db, err := strconv.Atoi(u.Path[1:])
	if err != nil {
		log.Printf(fmt.Errorf("parse redis db number: %w", err).Error())
		panic("")
	}

	return redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
}

func ConnectRedis() {
	RDB = newRDB()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func ClearRedis() {
	if !config.Dev {
		panic("ClearRedis only available in dev mode")
	}
	db := getRDB()
	keys, err := db.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	for _, key := range keys {
		db.Del(ctx, key)
	}
}

func handleRedisError(ctx string, err error) {
	if err != nil {
		log.Printf("redis error: %s: %s\n", ctx, err)
	}
}

func getRDB() *redis.Client {
	if RDB != nil {
		return RDB
	}
	return newRDB()
}

func Get(key string) string {
	value, err := getRDB().Get(ctx, key).Result()
	handleRedisError("Get "+key, err)
	return value
}

func GetObj(key string, obj any) {
	value, err := getRDB().Get(ctx, key).Bytes()
	handleRedisError("Get "+key, err)
	if err == nil {
		json.Unmarshal(value, obj)
	}
}

func SetObj(key string, obj any) {
	value, _ := json.Marshal(obj)
	_, err := getRDB().Set(ctx, key, value, 0).Result()
	handleRedisError("SetObj "+key, err)
}

func SetObjEx(key string, obj any, dur time.Duration) {
	value, _ := json.Marshal(obj)
	_, err := getRDB().Set(ctx, key, value, dur).Result()
	handleRedisError("SetObj "+key, err)
}

func Set(key string, value any) {
	_, err := getRDB().Set(ctx, key, value, 0).Result()
	handleRedisError("Set "+key, err)
}

func Set1(key string, value any) error {
	_, err := getRDB().Set(ctx, key, value, 0).Result()
	return err
}

func SetEx(key string, value any, dur time.Duration) {
	_, err := getRDB().Set(ctx, key, value, dur).Result()
	handleRedisError("SetEx "+key, err)
}

func RunSetEx(key string, value any, dur time.Duration) error {
	_, err := getRDB().Set(ctx, key, value, dur).Result()
	return err
}

func SetNx(key string, value any, dur time.Duration) bool {
	ok, err := getRDB().SetNX(ctx, key, value, dur).Result()
	handleRedisError("SetNx "+key, err)
	return ok
}

func Del(key string) {
	_, err := getRDB().Del(ctx, key).Result()
	handleRedisError("Del "+key, err)
}

func Expire(key string, dur time.Duration) {
	_, err := getRDB().Expire(ctx, key, dur).Result()
	handleRedisError("Expire "+key, err)
}

func Exists(key string) bool {
	ok, err := getRDB().Exists(ctx, key).Result()
	handleRedisError("Exists "+key, err)
	return ok == 1
}

func HSet(key string, values ...any) {
	_, err := getRDB().HSet(ctx, key, values).Result()
	handleRedisError("HSet "+key, err)
}

func HGetAll(key string) map[string]string {
	value, err := getRDB().HGetAll(ctx, key).Result()
	handleRedisError("HGetAll "+key, err)
	return value
}

func HGet(key, field string) string {
	value, err := getRDB().HGet(ctx, key, field).Result()
	handleRedisError("HGet "+key, err)
	return value
}

func HDel(key, field string) {
	_, err := getRDB().HDel(ctx, key, field).Result()
	handleRedisError("HDel "+key, err)
}

func HLen(key string) int64 {
	num, _ := getRDB().HLen(ctx, key).Result()
	return num
}

func SMembers(key string) []string {
	values, err := getRDB().SMembers(ctx, key).Result()
	handleRedisError("SMembers "+key, err)
	return values
}

func Keys(key string) []string {
	values, err := getRDB().Keys(ctx, key).Result()
	handleRedisError("Keys "+key, err)
	return values
}

func SPop(key string) string {
	value, err := getRDB().SPop(ctx, key).Result()
	handleRedisError("SPop "+key, err)
	return value
}

func SAdd(key string, members ...any) int64 {
	ret, err := getRDB().SAdd(ctx, key, members).Result()
	handleRedisError("SAdd "+key, err)
	return ret
}

func HExists(key, field string) bool {
	value, err := getRDB().HExists(ctx, key, field).Result()
	handleRedisError("HExists "+key, err)
	return value
}

func IncrBy(key string, value int64) int64 {
	ret, err := getRDB().IncrBy(ctx, key, value).Result()
	handleRedisError("IncrBy "+key, err)
	return ret
}

func Do(cmd ...any) (any, error) {
	return getRDB().Do(ctx, cmd...).Result()
}

func LPush(key string, value any) (int64, error) {
	n, err := getRDB().LPush(ctx, key, value).Result()
	handleRedisError("LPush "+key, err)
	return n, err
}

func RPush(key string, value any) error {
	_, err := getRDB().RPush(ctx, key, value).Result()
	handleRedisError("RPush "+key, err)
	return err
}

func RPop(key string, count int) error {
	_, err := getRDB().RPopCount(ctx, key, count).Result()
	handleRedisError("RPop "+key, err)
	return err
}

func RPop1(key string) string {
	v, err := getRDB().RPopCount(ctx, key, 1).Result()
	handleRedisError("RPop "+key, err)
	if err == nil && len(v) > 0 {
		return v[0]
	}
	return ""
}

func BLPop(key string, timeout time.Duration) (string, error) {
	result, err := getRDB().BLPop(ctx, timeout, key).Result()
	handleRedisError("BLPop "+key, err)
	if err != nil {
		return "", err
	}
	if len(result) == 2 {
		return result[1], nil
	}
	return "", fmt.Errorf("BLPop result: %#v", result)
}

func LRange(key string, start int, stop int) []string {
	values, err := getRDB().LRange(ctx, key, int64(start), int64(stop)).Result()
	handleRedisError("LRange "+key, err)
	return values
}
