package infrastructures

import (
	"github.com/go-redis/redis/v8"
)

type RedisHandler struct {
	Conn *redis.Client
}

func (r *RedisHandler) GetConn() *redis.Client {
	return r.Conn
}

func initRedis() *redis.Client {
	// TODO: options from env
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
