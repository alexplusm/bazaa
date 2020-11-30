package services

import (
	"fmt"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"time"
)

type AnswerService struct {
	AnswerRepo interfaces.IAnswerRepository
}

func (service *AnswerService) GetUserStatistics(
	userID string, totalOnly bool, games []bo.GameBO, from, to time.Time,
) error {
	gameIds := make([]string, 0, len(games))

	for _, game := range games {
		gameIds = append(gameIds, game.GameID)
	}

	res, err := service.AnswerRepo.SelectAnswersByUser(userID, gameIds, from, to)
	if err != nil {
		fmt.Println("error: ", err)
		return fmt.Errorf("get user statistics: %v", err)
	}

	fmt.Printf("Res: %+v | %v\n", res, len(res))

	return nil
}
