package services

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type RedisService struct {
	RedisClient interfaces.IRedisHandler
}

func (r *RedisService) Method() {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	fmt.Println("REDIS METHOD")

	err := conn.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := conn.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := conn.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
