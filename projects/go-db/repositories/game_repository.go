package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/models"
)

type GameRepository struct {
	DBConn interfaces.IDBHandler
}

func (repo *GameRepository) CreateGame(game models.GameModel) (string, error) {
	// conn := repo.DBConn.GetConn()
	// insert game in database and return gameID
	gameID := "kek-123/567"
	return gameID, nil
}

/* TODO: Examples */
//func (repository *PlayerRepository) GetPlayerByName(name string) (models.PlayerModel, error) {
//
//	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM player_models WHERE name = '%s'", name))
//	if err != nil {
//		return models.PlayerModel{}, err
//	}
//
//	var player models.PlayerModel
//
//	row.Next()
//	row.Scan(&player.Id, &player.Name, &player.Score)
//
//	return player, nil
//}

// TODO: use this later
//func (repository *PlayerRepositoryWithCircuitBreaker) GetPlayerByName(name string) (models.PlayerModel, error) {
//
//	output := make(chan models.PlayerModel, 1)
//	hystrix.ConfigureCommand("get_player_by_name", hystrix.CommandConfig{Timeout: 1000})
//	errors := hystrix.Go("get_player_by_name", func() error {
//
//		player, _ := repository.PlayerRepository.GetPlayerByName(name)
//
//		output <- player
//		return nil
//	}, nil)
//
//	select {
//	case out := <-output:
//		return out, nil
//	case err := <-errors:
//		println(err)
//		return models.PlayerModel{}, err
//	}
//}
