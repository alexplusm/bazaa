package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
)

type ScreenshotRepository struct {
	DBConn interfaces.IDBHandler
}
