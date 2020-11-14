package services

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

// TODO: rename RedisClient in all controllers
type ScreenshotCacheService struct {
	RedisClient interfaces.IRedisHandler
}

func (service *ScreenshotCacheService) GetScreenshot(
	gameID, userID string,
) (dao.ScreenshotURLDAO, bool) {
	// TODO: process errors
	ctx := context.Background()
	conn := service.RedisClient.GetConn()

	listLength, _ := conn.LLen(ctx, buildScreenshotsListKey(gameID)).Result()

	id, hasID := findScreenshot(conn, 0, listLength, gameID)
	if !hasID {
		return dao.ScreenshotURLDAO{}, false
	}

	url, _ := conn.HGet(ctx, id, screenshotURLKey).Result()
	service.setUserAnswerToScreenshot(userID, id, nullAnswerValue)

	return dao.ScreenshotURLDAO{ScreenshotID: id, ImageURL: url}, true
}

func (service *ScreenshotCacheService) setUserAnswerToScreenshot(
	userID, screenshotID, answer string,
) {
	ctx := context.Background()
	conn := service.RedisClient.GetConn()

	// TODO: research: HSet, HMSet, HSetNX
	err := conn.HSet(ctx, screenshotID, userID, answer).Err()
	if err != nil {
		fmt.Println(err) // TODO: process error
	}
}

func findScreenshot(conn *redis.Client, index, maxIndex int64, gameID string) (string, bool) {
	if index == maxIndex {
		return "", false
	}
	ctx := context.Background()

	id, _ := conn.LIndex(ctx, buildScreenshotsListKey(gameID), index).Result()
	keys, _ := conn.HKeys(ctx, id).Result()
	if len(keys) < consts.RequiredAnswerCountToFinishScreenshot+nonAnswerFieldsCount {
		return id, true
	} else {
		return findScreenshot(conn, index+1, maxIndex, gameID)
	}
}

//----------

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
		// TODO: incr index to find not full answered screenshot
		if index+1 == listLength {
			hasID = false
			id = ""
			break
		}
		id, err = conn.LIndex(ctx, buildScreenshotsListKey(gameID), index).Result()

		if err != nil {
			fmt.Println(err)
		}
		flag, err = service.screenshotNotHaveEnoughAnswers(id)
		index++
	}

	fmt.Println(err) // TODO: process error

	return hasID, id
}

func (service *ScreenshotCacheService) screenshotNotHaveEnoughAnswers(
	screenshotID string,
) (bool, error) {
	conn := service.RedisClient.GetConn()
	keys, err := conn.HKeys(context.Background(), screenshotID).Result()
	if err != nil {
		return false, fmt.Errorf("screenshot not have enough answers: %v", err)
	}
	fmt.Println("Keys:", keys)

	// INFO: проверка длинны массива ключей ответов с учетом наличия служебных полей
	maxKeysCount := consts.RequiredAnswerCountToFinishScreenshot + nonAnswerFieldsCount
	return len(keys) < maxKeysCount, nil // TODO: concurrency
}
