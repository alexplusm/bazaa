package services

import (
	"context"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ActiveUsersService struct {
	RedisClient     interfaces.IRedisHandler
	CacheKeyService interfaces.ICacheKeyService
}

func (service *ActiveUsersService) SetUserActivity(gameId, userId string) {
	conn := service.RedisClient.GetConn()
	key := service.CacheKeyService.GetActiveUserKey(gameId, userId)
	conn.Set(context.Background(), key, "1", time.Minute*15)
}

func (service *ActiveUsersService) CountOfActiveUsers(gameId string) (int, error) {
	conn := service.RedisClient.GetConn()
	pattern := service.CacheKeyService.GetActiveUserKeyPattern(gameId)
	keys, _ := conn.Keys(context.Background(), pattern).Result()
	return len(keys), nil
}
