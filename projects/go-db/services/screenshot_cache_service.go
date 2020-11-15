package services

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
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
	service.SetUserAnswerToScreenshot(userID, id, initAnswerValue)

	return dao.ScreenshotURLDAO{ScreenshotID: id, ImageURL: url}, true
}

func (service *ScreenshotCacheService) CanSetUserAnswerToScreenshot(
	userID, screenshotID string,
) bool {
	// TODO: process errors
	conn := service.RedisClient.GetConn()
	answer, _ := conn.HGet(context.Background(), screenshotID, userID).Result()

	return answer == initAnswerValue
}

func (service *ScreenshotCacheService) SetUserAnswerToScreenshot(
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

func (service *ScreenshotCacheService) ScreenshotExist(screenshotID string) bool {
	// TODO: process error in func
	conn := service.RedisClient.GetConn()
	res, _ := conn.HExists(
		context.Background(), screenshotID, screenshotURLKey,
	).Result()

	return res
}

func (service *ScreenshotCacheService) GetUsersAnswers(screenshotID string) []bo.UserAnswerCacheBO {
	// TODO: process error in func
	ctx := context.Background()
	conn := service.RedisClient.GetConn()

	// TODO: check screenshot exist
	keys, _ := conn.HKeys(ctx, screenshotID).Result()
	answers := make([]bo.UserAnswerCacheBO, 0, consts.RequiredAnswerCountToFinishScreenshot)

	for _, key := range keys {
		if serviceKeyMap[key] {
			continue
		}
		userAnswer, _ := conn.HGet(ctx, screenshotID, key).Result()
		if userAnswer != initAnswerValue {
			answer := bo.UserAnswerCacheBO{UserID: key, Answer: userAnswer}
			answers = append(answers, answer)
		}
	}

	return answers
}

// TODO: rename: на который не достаточно ответов
func findScreenshot(conn *redis.Client, index, maxIndex int64, gameID string) (string, bool) {
	if index == maxIndex {
		return "", false
	}
	ctx := context.Background()

	id, _ := conn.LIndex(ctx, buildScreenshotsListKey(gameID), index).Result()
	keys, _ := conn.HKeys(ctx, id).Result()
	// TODO: consts.RequiredAnswerCountToFinishScreenshot+nonAnswerFieldsCount -> in variable
	if len(keys) < consts.RequiredAnswerCountToFinishScreenshot+nonAnswerFieldsCount {
		return id, true
	} else {
		return findScreenshot(conn, index+1, maxIndex, gameID)
	}
}