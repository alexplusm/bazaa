package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type SourceRepository struct {
	DBConn interfaces.IDBHandler
}

const (
	insertSourceStatement = ``
)

func (repo *SourceRepository) InsertSource(source dao.SourceDAO) (string, error) {
	sourceID := ""
	return sourceID, nil
}
