package repos

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type AnswerRepo struct {
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
	selectAnsweredScreenshotsCountStatement = `
select COUNT(DISTINCT screenshot_id) FROM answers
WHERE answers.game_id = ($1)
`
	selectUniqueUsersInGameStatement = `
SELECT DISTINCT user_id FROM answers
WHERE answers.game_id = ($1)
`
	selectAnswersByGame = `
SELECT ans.user_id, ans.value, s.expert_answer, s.users_answer
FROM answers ans
INNER JOIN screenshots s
ON s.screenshot_id = ans.screenshot_id
WHERE
ans.game_id = ($1) AND
(ans.answer_date BETWEEN ($2) AND ($3))
`
	selectAnswersByUserAndGameStatement = `
SELECT ans.game_id, ans.screenshot_id, ans.answer_date, ans.value, s.expert_answer, s.users_answer
FROM answers ans
INNER JOIN screenshots s
ON s.screenshot_id = ans.screenshot_id
WHERE
ans.user_id = ($1) AND
ans.game_id = ($2) AND
(ans.answer_date BETWEEN ($3) AND ($4))
`
)

func (repo *AnswerRepo) InsertList(answers []dao.AnswerDAO) {
	for _, answer := range answers {
		err := repo.InsertOne(answer)
		if err != nil {
			fmt.Println("err: insert answers: ", err) // TODO: log error | return error
		}
	}
}

func (repo *AnswerRepo) InsertOne(answer dao.AnswerDAO) error {
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

func (repo *AnswerRepo) SelectScreenshotResult(gameID, screenshotID string) ([]dao.ScreenshotResultDAO, error) {
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

// TODO: check
func (repo *AnswerRepo) SelectAnsweredScreenshotsByGame(
	gameID string,
) (dao.AnsweredScreenshotsDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return dao.AnsweredScreenshotsDAO{},
			fmt.Errorf("select answered screenshots by game: acquire connection: %v", err)
	}
	defer conn.Release()

	var count int64

	row := conn.QueryRow(context.Background(), selectAnsweredScreenshotsCountStatement, gameID)
	if row.Scan(&count) != nil {
		return dao.AnsweredScreenshotsDAO{},
			fmt.Errorf("select answered screenshots by game: %v", err)
	}

	rows, err := conn.Query(context.Background(), selectUniqueUsersInGameStatement, gameID)
	if err != nil {
		return dao.AnsweredScreenshotsDAO{},
			fmt.Errorf("select answered screenshots by game: %v", err)
	}
	defer rows.Close()

	listUsers := make([]string, 0, 1024)

	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			log.Error("select answered screenshots by game: ", err)
			continue
		}
		listUsers = append(listUsers, user)
	}
	if rows.Err() != nil {
		log.Error("select answered screenshots by game: ", rows.Err())
	}

	res := dao.AnsweredScreenshotsDAO{
		Count:  int(count),
		UserID: listUsers,
	}

	fmt.Printf("reees: %+v\n", res)

	return res, nil
}

func (repo *AnswerRepo) SelectListTODO(gameID string, from, to time.Time) ([]dao.AnswerStatLeadDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("select answers by user: acquire connection: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(
		context.Background(), selectAnswersByGame,
		gameID, from.Unix(), to.Unix(),
	)

	if err != nil {
		return nil, fmt.Errorf("todooo: %v", err)
	}
	defer rows.Close()

	list := make([]dao.AnswerStatLeadDAO, 0, 1024)

	for rows.Next() {
		a := dao.AnswerStatLeadDAO{}
		var usersAnswer []byte

		err = rows.Scan(
			&a.UserID, &a.Value, &a.ExpertAnswer, &usersAnswer,
		)
		a.UsersAnswer = string(usersAnswer)
		if err != nil {
			log.Error("seletoodoo user: retrieve answer: ", err)
			continue
		}
		list = append(list, a)
	}
	if rows.Err() != nil {
		log.Error("toodoooooo: ", rows.Err())
	}

	return list, nil
}

func (repo *AnswerRepo) SelectListByUserAndGame(
	userID string, gameID string, from, to time.Time,
) ([]dao.UserAnswerDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("select answers by user and game: acquire connection: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(
		context.Background(), selectAnswersByUserAndGameStatement,
		userID, gameID, from.Unix(), to.Unix(),
	)

	if err != nil {
		return nil, fmt.Errorf("select answers by user and game: %v", err)
	}
	defer rows.Close()

	list := make([]dao.UserAnswerDAO, 0, 1024)

	for rows.Next() {
		a := dao.UserAnswerDAO{}
		var usersAnswer []byte

		err = rows.Scan(
			&a.GameID, &a.ScreenshotID, &a.AnswerDate,
			&a.Value, &a.ExpertAnswer, &usersAnswer,
		)
		a.UsersAnswer = string(usersAnswer)
		if err != nil {
			log.Error("select answers by user and game: retrieve answer: ", err)
			continue
		}
		list = append(list, a)
	}
	if rows.Err() != nil {
		log.Error("select answers by user and game: ", rows.Err())
	}

	return list, nil
}
