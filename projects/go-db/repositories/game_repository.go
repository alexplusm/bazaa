package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type GameRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertGameStatement = `
INSERT INTO games ("name", "start_date", "end_date", "answer_type", "question", "options_csv")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "game_id";
`
	insertGameWithoutOptionsStatement = `
INSERT INTO games ("name", "start_date", "end_date", "answer_type", "question")
VALUES ($1, $2, $3, $4, $5)
RETURNING "game_id";
`
	selectNotStartedGameCountWithSameIDStatement = `
SELECT COUNT(1) FROM games
WHERE "game_id" = $1 and "start_date" > $2;
`
)

func (repo *GameRepository) InsertGame(game dao.GameDAO) (string, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return "", fmt.Errorf("create game: acquire connection: %v", err)
	}
	defer conn.Release()

	var args []interface{}
	var statement string

	if game.AnswerType == consts.CategoricalAnswerType {
		args = []interface{}{
			game.Name, game.StartDate, game.EndDate, game.AnswerType, game.Question, game.Options,
		}
		statement = insertGameStatement
	} else {
		args = []interface{}{
			game.Name, game.StartDate, game.EndDate, game.AnswerType, game.Question,
		}
		statement = insertGameWithoutOptionsStatement
	}

	row := conn.QueryRow(context.Background(), statement, args...)

	var gameID string

	err = row.Scan(&gameID)
	if err != nil {
		return "", fmt.Errorf("create game: query: %v", err)
	}

	return gameID, nil // TODO:log: saving game into DB
}

func (repo *GameRepository) HasNotStartedGameWithSameID(gameID string) (bool, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return false, fmt.Errorf("has game: acquire connection: %v", err)
	}
	defer conn.Release()

	var gameCount int64
	now := time.Now().Unix()
	row := conn.QueryRow(context.Background(), selectNotStartedGameCountWithSameIDStatement, gameID, now)

	err = row.Scan(&gameCount)
	if err != nil {
		return false, fmt.Errorf("has game: query: %v", err)
	}

	return gameCount == 1, nil
}
