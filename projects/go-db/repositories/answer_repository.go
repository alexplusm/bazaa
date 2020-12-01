package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type AnswerRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertAnswerStatement = `
INSERT INTO answers ("screenshot_id", "game_id", "user_id", "value", "answer_date")
VALUES ($1, $2, $3, $4, $5);
`
	selectScreenshotResultsStatement = `
SELECT ans.user_id, ans.value, s.users_answer
FROM answers ans
INNER JOIN screenshots s
ON s.screenshot_id = ans.screenshot_id
WHERE
ans.game_id = ($1) AND ans.screenshot_id = ($2) 
`
	selectAnswersByUserStatement = `
SELECT ans.game_id, ans.screenshot_id, ans.answer_date, ans.value, s.expert_answer, s.users_answer
FROM answers ans
INNER JOIN screenshots s
ON s.screenshot_id = ans.screenshot_id
WHERE
ans.user_id = ($1) AND
(ans.answer_date BETWEEN ($2) AND ($3)) AND
ans.game_id IN `
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

	ts := time.Now().Unix()
	row, err := conn.Query(
		context.Background(),
		insertAnswerStatement,
		answer.ScreenshotID, answer.GameID, answer.UserID, answer.Value, ts,
	)
	if err != nil {
		return fmt.Errorf("insert answer: %v", err)
	}
	row.Close()

	return nil
}

func (repo *AnswerRepository) SelectAnswersByUser(
	userID string, gameIDs []string, from, to time.Time,
) ([]dao.AnswerStatDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("select answers by user: acquire connection: %v", err)
	}
	defer conn.Release()

	// TODO: костыль : добавлены ковычки!
	gamesVal := strings.Join(gameIDs, "','")
	fmt.Println("GamesValue: ", gamesVal)

	// TODO: костыль : добавлены ковычки! и аппендиться список gameID
	statement := selectAnswersByUserStatement + "('" + gamesVal + "');"

	rows, err := conn.Query(
		context.Background(), statement,
		userID, from.Unix(), to.Unix(),
	)

	if err != nil {
		return nil, fmt.Errorf("select answers by user: %v", err)
	}
	defer rows.Close()

	list := make([]dao.AnswerStatDAO, 0, 1024)

	for rows.Next() {
		a := dao.AnswerStatDAO{}
		var usersAnswer []byte

		err = rows.Scan(
			&a.GameID, &a.ScreenshotID, &a.AnswerDate,
			&a.Value, &a.ExpertAnswer, &usersAnswer,
		)
		a.UsersAnswer = string(usersAnswer)
		if err != nil {
			log.Error("select answers by user: retrieve answer: ", err)
			continue
		}
		list = append(list, a)
	}
	if rows.Err() != nil {
		log.Error("select answers by user: ", rows.Err())
	}

	return list, nil
}

func (repo *AnswerRepository) SelectScreenshotResult(gameID, screenshotID string) ([]dao.ScreenshotResultDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("select screenshot result: acquire connection: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(
		context.Background(), selectScreenshotResultsStatement,
		gameID, screenshotID,
	)
	if err != nil {
		return nil, fmt.Errorf("select screenshot result: %v", err)
	}
	defer rows.Close()

	list := make([]dao.ScreenshotResultDAO, 0, 10)

	for rows.Next() {
		r := dao.ScreenshotResultDAO{}
		var usersAnswer []byte

		err = rows.Scan(&r.UserID, &r.Value, &usersAnswer)
		r.UsersAnswer = string(usersAnswer)

		list = append(list, r)
	}
	if rows.Err() != nil {
		log.Error("select screenshot result: ", rows.Err())
	}

	return list, nil
}
