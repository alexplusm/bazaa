package repositories

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/consts"
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type GameRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertGameStatement = `
INSERT INTO games ("ext_system_id", "name", "start_date", "end_date", "answer_type", "question", "options_csv")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING "game_id";
`
	insertGameWithoutOptionsStatement = `
INSERT INTO games ("ext_system_id", "name", "start_date", "end_date", "answer_type", "question")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "game_id";
`
	selectGameWithSameIDStatement = `
SELECT "ext_system_id", "name", "start_date", "end_date", "answer_type", "question", "options_csv"
FROM games
WHERE "game_id" = $1;
`
	selectGames = `
SELECT "ext_system_id", "name", "start_date", "end_date", "answer_type", "question", "options_csv"
FROM games;
`
)

func (repo *GameRepository) InsertGame(game dao.GameDAO) (string, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return "", fmt.Errorf("insert game: acquire connection: %v", err)
	}
	defer conn.Release()

	var args []interface{}
	var statement string

	if game.AnswerType == consts.CategoricalAnswerType {
		args = []interface{}{
			game.ExtSystemID, game.Name, game.StartDate, game.EndDate,
			game.AnswerType, game.Question, game.Options,
		}
		statement = insertGameStatement
	} else {
		args = []interface{}{
			game.ExtSystemID, game.Name, game.StartDate,
			game.EndDate, game.AnswerType, game.Question,
		}
		statement = insertGameWithoutOptionsStatement
	}

	row := conn.QueryRow(context.Background(), statement, args...)

	var gameID string

	err = row.Scan(&gameID)
	if err != nil {
		return "", fmt.Errorf("insert game: query: %v", err)
	}

	return gameID, nil // TODO:log: saving game into DB
}

func (repo *GameRepository) SelectGame(gameID string) (dao.GameDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return dao.GameDAO{}, fmt.Errorf("select game: acquire connection: %v", err)
	}
	defer conn.Release()

	g := new(dao.GameDAO)
	row := conn.QueryRow(context.Background(), selectGameWithSameIDStatement, gameID)
	err = row.Scan(
		&g.ExtSystemID, &g.Name, &g.StartDate,
		&g.EndDate, &g.AnswerType, &g.Question, &g.Options,
	)

	if err != nil {
		return dao.GameDAO{}, fmt.Errorf("select game: %v", err)
	}

	return *g, nil
}

func (repo *GameRepository) SelectGames() ([]dao.GameDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("select games: acquire connection: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), selectGames)
	if err != nil {
		return nil, fmt.Errorf("select games: %v", err)
	}
	defer rows.Close()

	list := make([]dao.GameDAO, 0, 1024)

	for rows.Next() {
		g := new(dao.GameDAO)
		err = rows.Scan(
			&g.ExtSystemID, &g.Name, &g.StartDate,
			&g.EndDate, &g.AnswerType, &g.Question, &g.Options,
		)
		if err != nil {
			log.Error("select games: retrieve game: ", err)
			continue
		}
		list = append(list, *g)
	}
	if rows.Err() != nil {
		log.Error("select games: ", rows.Err())
	}

	return list, nil
}
