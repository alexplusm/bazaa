package repositories

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type SourceRepository struct {
	DBConn interfaces.IDBHandler
}

func (repo *SourceRepository) CreateSource(source dao.SourceDAO) error {

	return nil
}
