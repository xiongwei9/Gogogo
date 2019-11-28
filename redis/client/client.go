package client

import (
	"github.com/go-redis/redis/v7"
	"log"
)

var RedisClient *redis.Client

func init() {
	log.Println("init redis client")
	RedisClient = newClient()
}

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	log.Println(pong, err)

	return client
}
