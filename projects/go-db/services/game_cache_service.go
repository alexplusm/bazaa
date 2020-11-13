package services

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type GameCacheService struct {
	RedisClient    interfaces.IRedisHandler
	ScreenshotRepo interfaces.IScreenshotRepository
}

func (service *GameCacheService) PrepareGame(gameID string) {
	// достаем последнюю игру (за один час до начала игры) | todo: а если несколько игр?!
	// достаем все скрины из этой игры | todo: замешивание из прошлой игры с таким же типом

	fmt.Println("PrepareGame run ...")
	conn := service.RedisClient.GetConn()
	ctx := context.Background()

	screenshots, err := service.ScreenshotRepo.SelectScreenshotsByGameID(gameID)

	y := make([]interface{}, len(screenshots))
	for i := range screenshots {
		y[i] = screenshots[i].ScreenshotID
	}

	len, err := conn.RPush(ctx, screenshotsKey, y...).Result()
	if err != nil {
		fmt.Println(err)
	}

	// TODO: --- for testing
	fmt.Println("Length:", len)
	values, err := conn.LRange(ctx, screenshotsKey, 0, len).Result()

	fmt.Println("Values:", values)
	fmt.Println(err) // TODO: process err in func
}

func (service *GameCacheService) GameWithSameExtSystemIDExist(gameID, extSystemID string) bool {
	conn := service.RedisClient.GetConn()

	id, err := conn.HGet(context.Background(), gameID, extSystemIDKey).Result()
	if err != nil {
		fmt.Println("Error", err) // TODO: process error or return error
		return false
	}

	return id == extSystemID
}
