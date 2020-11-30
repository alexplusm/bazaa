package dao

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type AnswerDAO struct {
	AnswerID     int64
	UserID       string
	ScreenshotID string
	GameID       string
	Value        string
}

func (dest *AnswerDAO) FromCacheBO(src bo.UserAnswerCacheBO, gameID, screenshotID string) {
	dest.GameID = gameID
	dest.ScreenshotID = screenshotID
	dest.UserID = src.UserID
	dest.Value = src.Answer
}

type AnswerStatDAO struct {
	GameID       string
	ScreenshotID string
	AnswerDate   int64
	ExpertAnswer string
	UsersAnswer  string
}
