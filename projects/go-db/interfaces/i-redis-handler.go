package interfaces

import (
	"github.com/go-redis/redis/v8"
)

type IRedisHandler interface {
	GetConn() *redis.Client
}
