package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type AnswerService struct {
	AnswerRepo interfaces.IAnswerRepository
}

func (service *AnswerService) GetUserStatistics(userID, gameID string) error {
	return nil
}
