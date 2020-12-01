package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type ScreenshotRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertScreenshotStatement = `
INSERT INTO screenshots ("game_id", "source_id", "filename")
VALUES ($1, $2, $3);
`
	insertScreenshotWithExpertAnswerStatement = `
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
)

func (repo *ScreenshotRepository) SelectScreenshotsByGameID(gameID string) ([]dao.ScreenshotDAOFull, error) {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("select screenshots: acquire connection: %v", err)
	}
	defer conn.Release()

	row, err := conn.Query(context.Background(), selectScreenshotsByGameID, gameID)
	if err != nil {
		fmt.Println(err)
	}

	var screenshotID, sourceID, filename string
	var expertAnswer, usersAnswer []byte
	results := make([]dao.ScreenshotDAOFull, 0, 100)

	for row.Next() {
		err := row.Scan(&screenshotID, &sourceID, &filename, &expertAnswer, &usersAnswer)
		if err != nil {
			log.Error("select screenshots by game id: ", err)
			continue
		}
		obj := dao.ScreenshotDAOFull{
			ScreenshotID: screenshotID,
			SourceID:     sourceID,
			GameID:       gameID,
			Filename:     filename,
			ExpertAnswer: string(expertAnswer),
			UsersAnswer:  string(usersAnswer),
		}
		results = append(results, obj)
	}

	return results, nil
}

func (repo *ScreenshotRepository) InsertScreenshots(screenshots []dao.ScreenshotDAO) error {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("insert screenshots: acquire connection: %v", err)
	}
	defer conn.Release()

	for _, screenshot := range screenshots {
		// TODO: in goroutine?
		err := insertScreenshot(conn, screenshot)
		if err != nil {
			// TODO:log error
			continue
		}
	}

	return nil
}

func (repo *ScreenshotRepository) InsertScreenshotsWithExpertAnswer(screenshots []dao.ScreenshotWithExpertAnswerDAO) error {
	p := repo.DBConn.GetPool()
	conn, err := p.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("insert screenshots with expert answer: acquire connection: %v", err)
	}
	defer conn.Release()

	for _, screenshot := range screenshots {
		// TODO: in goroutine?
		err := insertScreenshotWithExpertAnswer(conn, screenshot)
		if err != nil {
			// TODO:log error
			fmt.Println("Error: ", err)
			continue
		}
	}

	return nil
}

func insertScreenshot(conn *pgxpool.Conn, s dao.ScreenshotDAO) error {
	row, err := conn.Query(
		context.Background(),
		insertScreenshotStatement,
		s.GameID, s.SourceID, s.Filename,
	)
	if err != nil {
		return fmt.Errorf("insert screenshot: %v", err)
	}
	row.Close()

	return nil
}

func insertScreenshotWithExpertAnswer(conn *pgxpool.Conn, s dao.ScreenshotWithExpertAnswerDAO) error {
	row, err := conn.Query(
		context.Background(),
		insertScreenshotWithExpertAnswerStatement,
		s.GameID, s.SourceID, s.Filename, s.ExpertAnswer,
	)
	if err != nil {
		return fmt.Errorf("insert screenshot: %v", err)
	}
	row.Close()

	return nil
}

func (repo *ScreenshotRepository) UpdateScreenshotUsersAnswer(screenshotID, usersAnswer string) error {
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

func (repo *ScreenshotRepository) ScreenshotExist(screenshotID string) (bool, error) {
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
