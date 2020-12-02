package services

import (
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type DurationService struct{}

func (service *DurationService) GetDurationByGame(from, to string, game bo.GameBO) (time.Time, time.Time) {
	return time.Now(), time.Now()
}
