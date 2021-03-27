package util

import "github.com/go-redis/redis/v8"

var r *redis.Client

func ConnectRedis() *redis.Client {
	r = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return r
}

func GetRedis() *redis.Client {
	return r
}
