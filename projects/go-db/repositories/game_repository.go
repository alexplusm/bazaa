package repositories

import (
	"context"
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type GameRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	createGameStatement = `
INSERT INTO games ("name", "start_date", "end_date", "answer_type", "question", "options_csv")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "game_id";
`
	createGameWithoutOptionsStatement = `
INSERT INTO games ("name", "start_date", "end_date", "answer_type", "question")
VALUES ($1, $2, $3, $4, $5)
RETURNING "game_id";
`
)

func (repo *GameRepository) CreateGame(game dao.GameDAO) (string, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return "", fmt.Errorf("create game: acquire connection: %v", err)
	}
	defer conn.Release()

	var args []interface{}
	var statement string

	if game.AnswerType == consts.Categorical {
		args = []interface{}{
			game.Name, game.StartDate, game.EndDate, game.AnswerType, game.Question, game.Options,
		}
		statement = createGameStatement
	} else {
		args = []interface{}{
			game.Name, game.StartDate, game.EndDate, game.AnswerType, game.Question,
		}
		statement = createGameWithoutOptionsStatement
	}

	row := conn.QueryRow(context.Background(), statement, args...)

	var gameID string

	err = row.Scan(&gameID)
	if err != nil {
		return "", fmt.Errorf("create game: query: %v", err)
	}

	return gameID, nil // TODO:log: saving game into DB
}

/* TODO: Examples */
//func (repository *PlayerRepository) GetPlayerByName(name string) (dao.PlayerModel, error) {
//
//	row, err :=repository.Query(fmt.Sprintf("SELECT * FROM player_models WHERE name = '%s'", name))
//	if err != nil {
//		return dao.PlayerModel{}, err
//	}
//
//	var player dao.PlayerModel
//
//	row.Next()
//	row.Scan(&player.Id, &player.Name, &player.Score)
//
//	return player, nil
//}

// TODO: use this later
//func (repository *PlayerRepositoryWithCircuitBreaker) GetPlayerByName(name string) (dao.PlayerModel, error) {
//
//	output := make(chan dao.PlayerModel, 1)
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
//		return dao.PlayerModel{}, err
//	}
//}
