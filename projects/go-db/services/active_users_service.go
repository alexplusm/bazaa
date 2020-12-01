package services

import (
	"context"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ActiveUsersService struct {
	RedisClient interfaces.IRedisHandler
}

func (service *ActiveUsersService) SetUserActivity(gameID, userID string) {
	conn := service.RedisClient.GetConn()
	// TODO: util FUNC |
	key := "++" + gameID + ":" + userID
	conn.Set(context.Background(), key, "1", time.Minute*15)
}

func (service *ActiveUsersService) CountOfActiveUsers(gameID string) (int, error) {
	conn := service.RedisClient.GetConn()
	pattern := "++" + gameID + ":" + "*"
	keys, _ := conn.Keys(context.Background(), pattern).Result()
	return len(keys), nil
}
