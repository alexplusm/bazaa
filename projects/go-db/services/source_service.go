package services

import (
	"fmt"
	"time"

	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/utils/logutils"
)

type SourceService struct {
	SourceRepo interfaces.ISourceRepo
}

func (service *SourceService) Create(gameId, value string, sourceType int) (string, error) {
	source := dao.SourceInsertDAO{
		SourceBaseDAO: dao.SourceBaseDAO{
			GameID: gameId, Type: sourceType, CreatedAt: time.Now().Unix(), Value: value,
		},
	}

	return service.SourceRepo.InsertOne(source)
}

func (service *SourceService) ListByGame(gameID string) ([]bo.SourceBO, error) {
	listDAO, err := service.SourceRepo.SelectListByGame(gameID)
	if err != nil {
		return nil, fmt.Errorf("source list by game: %v", err)
	}

	list := make([]bo.SourceBO, 0, len(listDAO))
	for _, source := range listDAO {
		list = append(list, source.ToBO())
	}

	return list, nil
}

// TODO
// Work only for sourceGames with type == Category !!!
func (service *SourceService) GameHasSomeSourceGameId(gameId, sourceGameId string) (bool, error) {
	// TODO: service.ListByGame | refactor bo.SourceBO
	sources, err := service.SourceRepo.SelectListByGame(gameId)
	if err != nil {
		return false, fmt.Errorf("%v GameHasSomeSourceGameId %v", logutils.GetStructName(service), err)
	}

	for _, source := range sources {
		if source.Value == sourceGameId {
			return true, nil
		}
	}

	return false, nil
}
