package services

import (
	"strings"
)

// TODO: redis key generator service !!!
//		and use this keys for communicate this redis

const (
	extSystemIDKey   = "extSystemID"
	screenshotsKey   = "--screenshots"
	gameKey          = "--game"
	screenshotURLKey = "url"

	// INFO: количество служебных полей: "url"
	nonAnswerFieldsCount = 1

	// INFO: default answer value
	initAnswerValue = "null"
)

var serviceKeyMap map[string]bool

func init() {
	serviceKeyMap = make(map[string]bool)

	serviceKeyMap[screenshotURLKey] = true
}

type CacheKeyService struct {
}

func (service *CacheKeyService) GetActiveUserKey(gameId, userId string) string {
	return "++" + gameId + ":" + userId
}

func (service *CacheKeyService) GetActiveUserKeyPattern(gameId string) string {
	return "++" + gameId + ":" + "*"
}

func (service *CacheKeyService) GetScreenshotListKey(gameId string) string {
	return screenshotsKey + ":" + gameId
}

func buildScreenshotsListKey(gameID string) string {
	return strings.Join([]string{screenshotsKey, gameID}, ":")
}

func buildGameKey(gameID string) string {
	return strings.Join([]string{gameKey, gameID}, ":")
}
