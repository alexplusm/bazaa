package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

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
)

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
