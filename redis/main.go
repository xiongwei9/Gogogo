package main

import (
	"github.com/go-redis/redis/v7"
	"github.com/xiongwei9/Gogogo/redis/client"
	"log"
	"time"
)

func testRedis() {
	redisClient := client.GetClient()

	const key = "name"
	const anotherKey = "a_name_not_exist"

	err := redisClient.Set(key, "xiongwei", time.Second*60).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(key).Result()
	if err != nil {
		panic(err)
	}
	log.Println(key, val)

	val2, err := redisClient.Get(anotherKey).Result()
	if err == redis.Nil {
		log.Printf("`%s` does not exist\n", anotherKey)
	} else if err != nil {
		panic(err)
	} else {
		log.Println(anotherKey, val2)
	}
}

func main() {
	testRedis()
}
