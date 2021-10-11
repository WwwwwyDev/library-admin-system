package core

import (
	"github.com/go-redis/redis"
)

func NewRedis(addr string, password string,db int) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // password set
		DB:       db,       // use default DB
	})
	return client
}
