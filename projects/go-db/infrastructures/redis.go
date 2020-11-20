package infrastructures

import (
	"os"

	"github.com/go-redis/redis/v8"
)

type RedisHandler struct {
	Conn *redis.Client
}

func (r *RedisHandler) GetConn() *redis.Client {
	return r.Conn
}

func initRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	addr := host + ":" + port

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
