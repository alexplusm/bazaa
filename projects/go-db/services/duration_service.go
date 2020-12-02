package services

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/utils"
)

type DurationService struct{}

func (service *DurationService) GetDurationByGame(from, to string, game bo.GameBO) (time.Time, time.Time) {
	year, month, day := time.Now().Date()
	fromRes := game.StartDate
	toRes := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	if from != "" {
		parsedFrom, err := utils.FromTimestampToTime(from)
		if err != nil {
			// TODO: log?
		} else {
			fromRes = parsedFrom
		}
	}
	if to != "" {
		parsedTo, err := utils.FromTimestampToTime(to)
		if err != nil {
			// TODO: log?
		} else {
			toRes = parsedTo
		}
	}

	return fromRes, toRes
}
