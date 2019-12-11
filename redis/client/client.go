package client

import (
	"log"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func init() {
	log.Println("init redis client")
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis ping error: %v", err)
	}
}

func GetClient() *redis.Client {
	return client
}
