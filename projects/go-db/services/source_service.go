package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type SourceService struct {
	SourceRepo interfaces.ISourceRepository
}

func (repo *SourceService) ListByGame(gameID string) ([]dao.Source2DAO, error) {
	return repo.SourceRepo.SelectListByGame(gameID)
}
