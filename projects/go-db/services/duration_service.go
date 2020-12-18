package services

import (
	"fmt"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/timeutils"
)

type DurationService struct{}

func (service *DurationService) GetDurationByGame(from, to string, game bo.GameBO) (time.Time, time.Time, error) {
	year, month, day := time.Now().Date()
	fromResult := game.StartDate
	toResult := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	if from != "" {
		parsedFrom, err := timeutils.FromTimestampToTime(from)
		if err != nil {
			return fromResult, toResult, fmt.Errorf("get duration by game: %v", err)
		} else {
			fromResult = parsedFrom
		}
	}
	if to != "" {
		parsedTo, err := timeutils.FromTimestampToTime(to)
		if err != nil {
			return fromResult, toResult, fmt.Errorf("get duration by game: %v", err)
		} else {
			toResult = parsedTo
		}
	}

	return fromResult, toResult, nil
}
