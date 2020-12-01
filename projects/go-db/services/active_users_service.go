package services

import (
	"context"
	"fmt"
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

	fmt.Println("### KEEEEY: ", key)

	conn.Set(context.Background(), key, "1", time.Minute*60)
}

func (service *ActiveUsersService) CountOfActiveUsers(gameID string) (int, error) {
	conn := service.RedisClient.GetConn()

	pattern := "++" + gameID + ":" + "*"

	fmt.Println("Pattern: ", pattern)

	keys, _ := conn.Keys(context.Background(), pattern).Result()
	fmt.Println("KEEEEYS", keys)
	fmt.Println(len(keys))
	return len(keys), nil
}
