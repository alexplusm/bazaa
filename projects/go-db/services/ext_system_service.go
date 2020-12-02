package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type ExtSystemService struct {
	ExtSystemRepo interfaces.IExtSystemRepository
}

func (service *ExtSystemService) CreateExtSystem(extSystem bo.ExtSystemBO) (string, error) {
	extSystemDAO := dao.ExtSystemDAO{}
	extSystemDAO.FromBO(extSystem)

	return service.ExtSystemRepo.InsertExtSystem(extSystemDAO)
}

func (service *ExtSystemService) ExtSystemExist(extSystemID string) (bool, error) {
	return service.ExtSystemRepo.ExtSystemExist(extSystemID)
}

func (service *ExtSystemService) ExtSystemList() ([]bo.ExtSystemBO, error) {
	list, err := service.ExtSystemRepo.SelectExtSystems()
	if err != nil {
		return nil, err
	}

	listBO := make([]bo.ExtSystemBO, 0, len(list))
	for _, item := range list {
		listBO = append(listBO, item.ToBO())
	}
	return listBO, nil
}
