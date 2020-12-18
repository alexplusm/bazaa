package repos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/logutils"
)

type ScreenshotRepo struct {
	DBConn interfaces.IDBHandler
}

const (
	insertScreenshotStatement = `
INSERT INTO screenshots ("game_id", "source_id", "filename", "expert_answer")
VALUES ($1, $2, $3, $4);
`
	selectScreenshotsByGameID = `
SELECT "screenshot_id", "source_id", "filename", "expert_answer", "users_answer"
FROM screenshots
WHERE screenshots.game_id = $1
`
	updateScreenshotUsersAnswerStatement = `
UPDATE screenshots
SET users_answer = ($1)
WHERE screenshots.screenshot_id = ($2)
`
	existScreenshotStatement = `
SELECT COUNT(1)
FROM screenshots
WHERE "screenshot_id" = ($1);
`
	selectScreenshotCountByUserStatement = `
SELECT COUNT(*) FROM screenshots 
WHERE screenshots.game_id = ($1)
`
)

func (repo *ScreenshotRepo) SelectListByGameID(gameID string) ([]dao.ScreenshotRetrieveDAO, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("%v SelectListByGameID: acquire connection: %v", logutils.GetStructName(repo), err)
	}
	defer conn.Release()

	row, err := conn.Query(context.Background(), selectScreenshotsByGameID, gameID)
	if err != nil {
		return nil, fmt.Errorf("%v SelectListByGameID %v", logutils.GetStructName(repo), err)
	}

	result := make([]dao.ScreenshotRetrieveDAO, 0, 512)

	for row.Next() {
		s := dao.ScreenshotRetrieveDAO{GameID: gameID}

		err := row.Scan(&s.ScreenshotID, &s.SourceID, &s.Filename, &s.ExpertAnswer, &s.UsersAnswer)
		if err != nil {
			log.Error(logutils.GetStructName(repo), "SelectListByGameID: ", err)
			continue
		}

		result = append(result, s)
	}

	return result, nil
}

func (repo *ScreenshotRepo) InsertList(screenshots []dao.ScreenshotCreateDAO) error {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("insert screenshots: acquire connection: %v", err)
	}
	defer conn.Release()

	for _, screenshot := range screenshots {
		err := insertScreenshot(conn, screenshot)
		if err != nil {
			// TODO:log error
			continue
		}
	}

	return nil
}

func insertScreenshot(conn *pgxpool.Conn, s dao.ScreenshotCreateDAO) error {
	row, err := conn.Query(
		context.Background(),
		insertScreenshotStatement,
		s.GameID, s.SourceID, s.Filename, s.ExpertAnswer,
	)
	if err != nil {
		return fmt.Errorf("insert screenshot: %v", err)
	}
	row.Close()

	return nil
}

func (repo *ScreenshotRepo) UpdateUsersAnswer(screenshotID, usersAnswer string) error {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("update screenshot users answer: acquire connection: %v", err)
	}
	defer conn.Release()

	row, err := conn.Query(
		context.Background(), updateScreenshotUsersAnswerStatement,
		usersAnswer, screenshotID,
	)
	if err != nil {
		return fmt.Errorf("update screenshot users answer: %v", err)
	}
	defer row.Close()

	return nil
}

func (repo *ScreenshotRepo) Exist(screenshotID string) (bool, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return false, fmt.Errorf("screenshot exist: acquire connection: %v", err)
	}
	defer conn.Release()

	var count int64

	row := conn.QueryRow(context.Background(), existScreenshotStatement, screenshotID)
	if row.Scan(&count) != nil {
		return false, fmt.Errorf("screenshot exist: %v", err)
	}

	return count != 0, nil
}

func (repo *ScreenshotRepo) CountByGame(gameID string) (int, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return -1, fmt.Errorf("screenshot count by game: acquire connection: %v", err)
	}
	defer conn.Release()

	var count int64

	row := conn.QueryRow(context.Background(), selectScreenshotCountByUserStatement, gameID)
	if row.Scan(&count) != nil {
		return -1, fmt.Errorf("screenshot count by game: %v", err)
	}

	return int(count), nil
}
