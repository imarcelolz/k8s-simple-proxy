package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})

	if err := redisClient.Ping().Err(); err != nil {
		panic(err)
	}

	topic := redisClient.PSubscribe("*")
	channel := topic.Channel()

	for msg := range channel {
		fmt.Println(msg.Payload)
	}
}
