package cachemanager

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
)

type Cache struct {
	Redis *redis.Client
}

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
