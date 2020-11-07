package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type ScreenshotRepository struct {
	DBConn interfaces.IDBHandler
}

func (repo *ScreenshotRepository) InsertScreenshots(screenshots []dao.ScreenshotDAO) error {
	return nil
}

func (repo *ScreenshotRepository) InsertScreenshotsWithExpertAnswer(screenshots []dao.ScreenshotWithExpertAnswerDAO) error {
	return nil
}
