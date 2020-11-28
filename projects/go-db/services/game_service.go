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

func (service *GameService) GetGames(extSystemID string) ([]bo.GameBO, error) {
	// INFO: filter by "extSystemID" in SQL statement | for performance
	// remove filtering in this func
	gamesDAO, err := service.GameRepo.SelectGames()
	if err != nil {
		return nil, fmt.Errorf("get games: %v", err)
	}

	list := make([]bo.GameBO, 0, len(gamesDAO))

	for _, gameDAO := range gamesDAO {
		if gameDAO.ExtSystemID == extSystemID {
			list = append(list, gameDAO.ToBO())
		}
	}

	return list, nil
}
