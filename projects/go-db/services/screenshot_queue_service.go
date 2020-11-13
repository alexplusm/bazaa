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

func (r *RedisService) Method() {
	gameID := "bd255325-e7d1-44bd-8f76-1ff4796e71a2"

	//ctx := context.Background()
	//conn := r.RedisClient.GetConn()

	// CHECK GAME <-> EXTERNAL_SYSTEM_ID
	//sysID := "ExtSys-123-id"
	//e := r.checkGame(gameID, sysID)
	//fmt.Println("1) Has GAME!", e)
	//r.insertGame(gameID, sysID)
	//e = r.checkGame(gameID, sysID)
	//fmt.Println("2) Has GAME!", e)
	// CHECK GAME <-> EXTERNAL_SYSTEM_ID END

	// --------------------------------------
	//r.insertScreenshots(gameID, screenshots)
	//keys, err := conn.Keys(ctx, "*").Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("KEYS: ", keys)
	//fmt.Println("LEN KEYS:", len(keys))
	//fmt.Println("-----")
	//for _, key := range keys {
	//	v, err := conn.HGet(ctx, key, "url").Result()
	//	fmt.Println("###", key, "|", v, err)
	//}
	//fmt.Println(len(screenshots))
	// --------------------------------------

	hasID, id := r.getScreenshotID(gameID)
	fmt.Println("===", hasID, id)
}

// TODO: move to game_cache_service
func (r *RedisService) insertGame(gameID, externalSystemID string) {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	conn.HSet(ctx, gameID, extSystemIDKey, externalSystemID)
}

// скриншот, который можно отдать пользователю
func (r *RedisService) getScreenshotID(gameID string) (bool, string) {
	ctx := context.Background()
	conn := r.RedisClient.GetConn()

	listLength, err := conn.LLen(ctx, getScreenshotsListKey(gameID)).Result()
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
		id, err = conn.LIndex(ctx, getScreenshotsListKey(gameID), index).Result()

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
