package dao

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

// TODO: AnswerInsertDAO
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

// INFO:
//		use this struct for queries
//		which require JOIN on "answers" and "screenshots" tables
type AnswerScreenshotRetrieveDAO struct {
	GameID       string
	UserID       string
	ScreenshotID string
	AnswerDate   int64
	Value        string
	UsersAnswer  []byte
	ExpertAnswer []byte
}
