package db

import (
	"dreamland/pkg"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
	"time"
)

// client redis链接池
var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password:    viper.GetString("redis.password"),
		DB:          viper.GetInt("redis.db"),
		PoolSize:    viper.GetInt("redis.pool"),
		MaxRetries:  3,
		IdleTimeout: time.Second * 10,
	})

	pong, err := client.Ping().Result()
	if err == redis.Nil {
		panic(pkg.GetMsg(pkg.RedisError))
	} else if err != nil {
		panic(err)
	} else {
		log.Println(pong)
	}
}

func Exist(key string) bool {
	_, err := client.Exists(key).Result()
	if err != nil {
		return false
	}
	return true
}

func Set(key string, value interface{}, expiration time.Duration) bool {
	_, err := client.Set(key, value, expiration).Result()
	if err != nil {
		return false
	}
	return true
}

func Get(key string) (string, error) {
	return client.Get(key).Result()
}

func HExists(key, field string) bool {
	ok, err := client.HExists(key, field).Result()
	if err != nil {
		return false
	}
	return ok
}

// HashSet 向key的hash中添加元素field的值
func HashSet(key, field string, data interface{}) (bool, error) {
	return client.HSet(key, field, data).Result()
}

// BatchHashSet 批量向key的hash添加对应元素field的值
func BatchHashSet(key string, fields map[string]interface{}) (string, error) {
	return client.HMSet(key, fields).Result()
}

// HashGet 通过key获取hash的元素值
func HashGet(key, field string) (string, error) {
	val, err := client.HGet(key, field).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// BatchHashGet 批量获取key的hash中对应多元素值
func BatchHashGet(key string, fields ...string) ([]interface{}, error) {
	return client.HMGet(key, fields...).Result()
}
