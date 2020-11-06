package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type ScreenshotRepository struct {
	DBConn interfaces.IDBHandler
}

func (repo *SourceRepository) InsertScreenshot(screenshot dao.ScreenshotDAO) error {
	return nil
}
