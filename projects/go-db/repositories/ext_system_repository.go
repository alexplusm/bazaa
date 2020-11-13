package repositories

import "github.com/Alexplusm/bazaa/projects/go-db/interfaces"

type ExtSystemRepository struct {
	DBConn interfaces.IDBHandler
}

func (repo *ExtSystemRepository) InsertExtSystem() error {
	return nil
}
