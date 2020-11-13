package services

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ScreenshotCacheService struct {
	RedisClient interfaces.IRedisHandler
}

// скриншот, который можно отдать пользователю
func (service *ScreenshotCacheService) getScreenshotID(gameID string) (bool, string) {
	ctx := context.Background()
	conn := service.RedisClient.GetConn()

	listLength, err := conn.LLen(ctx, buildScreenshotsListKey(gameID)).Result()
	var index int64
	var id string
	var hasID = true
	flag := true

	for flag {
		// todo: incr index to find not full answered screenshot
		if index+1 == listLength {
			hasID = false
			id = ""
			break
		}
		id, err = conn.LIndex(ctx, buildScreenshotsListKey(gameID), index).Result()

		if err != nil {
			fmt.Println(err)
		}
		flag = service.checkAnsweredScreenshot(id)
		index++
	}

	fmt.Println(err) // TODO: process error

	return hasID, id
}

func (service *ScreenshotCacheService) checkAnsweredScreenshot(screenshotID string) bool {
	ctx := context.Background()
	conn := service.RedisClient.GetConn()

	keys, err := conn.HKeys(ctx, screenshotID).Result()
	fmt.Println("Keys:", keys)

	fmt.Println(err) // TODO: process error
	// INFO: проверка длинны массива ключей ответов с учетом наличия служебных полей
	maxKeysCount := consts.RequiredAnswerCountToFinishScreenshot + nonAnswerFieldsCount
	return len(keys) < maxKeysCount // TODO: concurrency
}
