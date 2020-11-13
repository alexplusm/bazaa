package services

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

// TODO: move to consts
const (
	extSystemIDKey   = "extSystemID"
	screenshotsKey   = "__screenshots"
	gameKey          = "__game"
	screenshotURLKey = "url"
	answerCount      = 5 // TODO: rename + move to config
	// INFO: "url" ...
	serviceFields = 1
)

type RedisService struct {
	RedisClient interfaces.IRedisHandler
}

// скриншот, который можно отдать пользователю
func (r *RedisService) getScreenshotID(gameID string) (bool, string) {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

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
		flag = r.checkAnsweredScreenshot(id)
		index++
	}

	fmt.Println(err) // TODO: process error

	return hasID, id
}

func (r *RedisService) checkAnsweredScreenshot(screenshotID string) bool {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	keys, err := conn.HKeys(ctx, screenshotID).Result()
	fmt.Println("Keys:", keys)

	fmt.Println(err) // TODO: process error
	// INFO: проверка длинны массива ключей ответов с учетом наличия поля служебных полей
	return len(keys) < answerCount+serviceFields // TODO: concurrency
}
