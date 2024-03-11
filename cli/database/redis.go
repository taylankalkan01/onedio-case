package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnectWithRedis() *redis.Client {
	addr := "localhost:6379"
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Redis connection error: ", err)
	}
	return client
}
