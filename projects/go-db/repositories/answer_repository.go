package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type AnswerRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertAnswerStatement = `
INSERT INTO answers ("screenshot_id", "game_id", "user_id", "value")
VALUES ($1, $2, $3, $4);
`
	selectAnswersByUser = `
SELECT "game_id", "screenshot_id", "value"
FROM answers
WHERE "user_id" = ($1) && "answer_date" >= ($2) && "answer_date" <= ($3);
`
)

func (repo *AnswerRepository) InsertAnswers(answers []dao.AnswerDAO) {
	for _, answer := range answers {
		err := repo.InsertAnswer(answer)
		if err != nil {
			fmt.Println("err: insert answers: ", err) // TODO: log error | return error
		}
	}
}

func (repo *AnswerRepository) InsertAnswer(answer dao.AnswerDAO) error {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("insert game: acquire connection: %v", err)
	}
	defer conn.Release()

	row, err := conn.Query(
		context.Background(),
		insertAnswerStatement,
		answer.ScreenshotID, answer.GameID, answer.UserID, answer.Value,
	)
	if err != nil {
		return fmt.Errorf("insert answer: %v", err)
	}
	row.Close()

	return nil
}

func (repo *AnswerRepository) SelectAnswersByUser(userID string, from, to time.Time) error {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("select answers by user: acquire connection: %v", err)
	}
	defer conn.Release()

	//from.Year()
	//from.Month()
	//from.Day()

	rows, err := conn.Query(context.Background(), selectAnswersByUser, userID)
	if err != nil {
		return fmt.Errorf("select games: %v", err)
	}
	defer rows.Close()

	/*
		rows, err := conn.Query(context.Background(), selectGames)
		if err != nil {
			return nil, fmt.Errorf("select games: %v", err)
		}
		defer rows.Close()

		list := make([]dao.GameDAO, 0, 1024)

		for rows.Next() {
			g := new(dao.GameDAO)
			err = rows.Scan(
				&g.GameID, &g.ExtSystemID, &g.Name, &g.StartDate,
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


	*/

	return nil
}
