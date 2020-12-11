package services

import (
	"fmt"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
)

type SourceService struct {
	SourceRepo interfaces.ISourceRepo
}

func (repo *SourceService) ListByGame(gameID string) ([]bo.SourceBO, error) {
	listDAO, err := repo.SourceRepo.SelectListByGame(gameID)
	if err != nil {
		return nil, fmt.Errorf("source list by game: %v", err)
	}

	list := make([]bo.SourceBO, 0, len(listDAO))
	for _, source := range listDAO {
		list = append(list, source.ToBO())
	}

	return list, nil
}
