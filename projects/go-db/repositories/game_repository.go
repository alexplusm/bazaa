package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type GameRepository struct {
	interfaces.IDBHandler
}

func (repo *GameRepository) CreateGame() (string, error) {
	// conn := repo.GetConn()
	// insert game in database and return gameID
	gameID := ""
	return gameID, nil
}

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
