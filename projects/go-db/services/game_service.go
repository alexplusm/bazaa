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

func (service *GameService) Create(game bo.GameBO) (string, error) {
	gameDAO := dao.GameDAO{}
	gameDAO.FromBO(game)

	return service.GameRepo.InsertOne(gameDAO)
}

func (service *GameService) Retrieve(gameID string) (bo.GameBO, error) {
	gameDAO, err := service.GameRepo.SelectOne(gameID)
	if err != nil {
		return bo.GameBO{}, fmt.Errorf("get game: %v", err)
	}

	gameBO := gameDAO.ToBO()

	return gameBO, nil
}

func (service *GameService) List(extSystemID string) ([]bo.GameBO, error) {
	gamesDAO, err := service.GameRepo.SelectList(extSystemID)
	if err != nil {
		return nil, fmt.Errorf("get games: %v", err)
	}

	list := make([]bo.GameBO, 0, len(gamesDAO))

	for _, gameDAO := range gamesDAO {
		list = append(list, gameDAO.ToBO())
	}

	return list, nil
}

func (service *GameService) Exist(gameID string) (bool, error) {
	return service.GameRepo.Exist(gameID)
}

func (service *GameService) FilterGames(gamesID []string, games []bo.GameBO) []bo.GameBO {
	if len(gamesID) == 0 {
		return games
	}

	filteredGames := make([]bo.GameBO, 0, len(gamesID))

	for _, game := range games {
		for _, gameID := range gamesID {
			if game.GameID == gameID {
				filteredGames = append(filteredGames, game)
			}
		}
	}

	return filteredGames
}

func (service *GameService) GetEarliestGame(games []bo.GameBO) bo.GameBO {
	result := games[0]

	for _, game := range games {
		if result.StartDate.Before(game.StartDate) {
			result = game
		}
	}

	return result
}
