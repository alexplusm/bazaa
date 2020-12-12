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

// AnswerScreenshotRetrieveDAO
// INFO:
//		use this struct for queries
//		which require JOIN on "answers" and "screenshots" tables
type UserAnswerDAO struct {
	GameID       string
	ScreenshotID string
	AnswerDate   int64

	Value        string
	ExpertAnswer string
	UsersAnswer  string
}

// AnswerRetrieveDAO

type AnswerRetrieve2DAO struct {
	UserID       string
	Value        string
	UsersAnswer  []byte
	ExpertAnswer []byte
}
