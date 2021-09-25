package common

import "github.com/go-redis/redis/v8"

var R *redis.Client

func ConnectRedis() *redis.Client {
	R = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return R
}

func GetRedis() *redis.Client {
	return R
}