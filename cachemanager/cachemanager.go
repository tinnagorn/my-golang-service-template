package cachemanager

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Cache struct {
	Redis *redis.Client
}

var RedisClientConn *Cache

func NewCache(redisHost string, redisPort string, redisPassword string, redisDB int) (*Cache, error) {
	var redisClient *redis.Client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDB,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println("Connect to Redis " + redisHost + ":" + redisPort)
		return nil, errors.New("Cannot connect to redis. " + err.Error())
	}

	return &Cache{
		Redis: redisClient,
	}, nil
}

func (cache *Cache) Close() error {
	return cache.Close()
}

func (cache *Cache) GetCache(key string) (result string, err error) {

	value, err := cache.Redis.Get(key).Result()
	if err == redis.Nil {
		return "", redis.Nil
	} else if err != nil {
		return "", errors.Wrap(err, "Unable to get value by key")
	}
	return value, nil
}

func (cache *Cache) SetCache(key string, value string, expire time.Duration) error {
	err := cache.Redis.Set(key, value, expire).Err()
	if err != nil {
		return errors.Wrap(err, "Unable to set value")
	}
	return nil
}

func (cache *Cache) DeleteCache(key string) error {
	err := cache.Redis.Del(key).Err()
	if err != nil {
		return errors.Wrap(err, "Unable to delete value")
	}
	return nil
}

func NewRedisClient() error {
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetString("redis.port")
	redisPw := viper.GetString("secrets.redis.password")
	redisDb := viper.GetInt("redis.db")
	client, err := NewCache(redisHost, redisPort, redisPw, redisDb)
	if err != nil {
		log.Fatalf("Error on init Redis %s", err.Error())
		return err
	}
	RedisClientConn = client
	return nil
}
func GetRedisClient() *Cache {
	return RedisClientConn
}

func Close() {
	RedisClientConn.Redis.Close()
}

func GetCacheOutput(key string, output interface{}) (err error) {
	// Get Data from Redis
	val1, err := RedisClientConn.GetCache(key)
	if err != nil {
		return err
	}

	// Unmarshal string and return
	return json.Unmarshal([]byte(val1), output)

}

func SetCacheOutPut(key string, input interface{}) (err error) {

	// Convert data to JSON
	serialized, err := json.Marshal(input)

	if err != nil {
		return err
	}
	// Set to Redis with duration 1 hour
	err = RedisClientConn.SetCache(key, string(serialized), 1*time.Hour)
	if err != nil {
		return err
	}

	return nil
}
