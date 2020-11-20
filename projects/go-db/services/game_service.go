package services

import (
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type GameService struct {
	GameRepo interfaces.IGameRepository
}

func (service *GameService) CreateGame(game bo.GameBO) (string, error) {
	gameDAO := dao.GameDAO{}
	gameDAO.FromBO(game)

	return service.GameRepo.InsertGame(gameDAO)
}

func (service *GameService) GetGame(gameID string) (bo.GameBO, error) {
	gameDAO, err := service.GameRepo.SelectGame(gameID)
	if err != nil {
		return bo.GameBO{}, fmt.Errorf("get game: %v", err)
	}

	gameBO := gameDAO.ToBO()

	return gameBO, nil
}
