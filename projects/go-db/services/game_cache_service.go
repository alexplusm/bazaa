package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type GameCacheService struct {
	RedisClient    interfaces.IRedisHandler
	ScreenshotRepo interfaces.IScreenshotRepository
	GameRepo       interfaces.IGameRepository
}

// TODO: достаем последнюю игру (за один час до начала игры, например)
func (service *GameCacheService) PrepareGame(gameID string) error {
	game, err := service.GameRepo.SelectGame(gameID)
	if err != nil {
		return fmt.Errorf("prepare game: %v", err)
	}

	fmt.Println("ExtID", game.ExtSystemID)

	err = service.insertGame(gameID, game.ExtSystemID)
	if err != nil {
		return fmt.Errorf("prepare game: %v", err)
	}
	err = service.insertScreenshots(gameID)
	if err != nil {
		return fmt.Errorf("prepare game: %v", err)
	}
	return nil
}

func (service *GameCacheService) GameWithSameExtSystemIDExist(gameID, extSystemID string) bool {
	conn := service.RedisClient.GetConn()

	id, err := conn.HGet(context.Background(), gameID, extSystemIDKey).Result()
	if err != nil {
		// TODO: process error or return error
		fmt.Println("GameWithSameExtSystemIDExist: ", err)
		return false
	}

	return id == extSystemID
}

func (service *GameCacheService) insertGame(gameID, extSystemID string) error {
	conn := service.RedisClient.GetConn()
	return conn.HSet(
		context.Background(), buildGameKey(gameID), extSystemIDKey, extSystemID,
	).Err()
}

func (service *GameCacheService) insertScreenshots(gameID string) error {
	// TODO: замешивание из прошлой игры с таким же типом
	conn := service.RedisClient.GetConn()
	ctx := context.Background()
	key := buildScreenshotsListKey(gameID)

	lengthInCache, err := conn.LLen(ctx, key).Result()
	cachedIDs, err := conn.LRange(ctx, key, 0, lengthInCache).Result()
	screenshots, err := service.ScreenshotRepo.SelectScreenshotsByGameID(gameID)
	mergedScreenshots := mergeScreenshotsWithCache(cachedIDs, screenshots)

	if len(mergedScreenshots) > 0 {
		list, idURLMap := convertToInterfaces(mergedScreenshots)

		for id, url := range idURLMap {
			err = conn.HSet(ctx, id, screenshotURLKey, url).Err()
		}
		err = conn.RPush(ctx, key, list...).Err()
	}

	fmt.Println(err)

	return nil // TODO: process err in func
}

func buildScreenshotsListKey(gameID string) string {
	return strings.Join([]string{screenshotsKey, gameID}, ":")
}

func buildGameKey(gameID string) string {
	return strings.Join([]string{gameKey, gameID}, ":")
}

func mergeScreenshotsWithCache(cache []string, screenshots []dao.ScreenshotDAOFull) []dao.ScreenshotDAOFull {
	screenshotCachedMap := make(map[string]bool)
	newScreenshots := make([]dao.ScreenshotDAOFull, 0, len(screenshots))

	for _, screenshotIDInCache := range cache {
		screenshotCachedMap[screenshotIDInCache] = true
	}
	for i := range screenshots {
		// INFO: screenshotIDs in cache must be unique
		if !screenshotCachedMap[screenshots[i].ScreenshotID] {
			newScreenshots = append(newScreenshots, screenshots[i])
		}
	}

	return newScreenshots
}

func convertToInterfaces(screenshots []dao.ScreenshotDAOFull) ([]interface{}, map[string]interface{}) {
	resultList := make([]interface{}, 0, len(screenshots))
	resultMap := make(map[string]interface{})
	for _, screenshot := range screenshots {
		resultList = append(resultList, screenshot.ScreenshotID)
		resultMap[screenshot.ScreenshotID] = buildFileURL(screenshot.Filename)
	}

	return resultList, resultMap
}

// TODO: move + add logic ?
func buildFileURL(filename string) string {
	return filename
}
