package redis

import "github.com/go-redis/redis/v8"

var RedisClient *redis.Client

func SimpleInit(host string, password string, db int) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
	defer RedisClient.Close()
}

func AllFiledInit(options *redis.Options){
	RedisClient := redis.NewClient(options)
	defer RedisClient.Close()
}