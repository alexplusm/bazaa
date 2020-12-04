package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type ExtSystemService struct {
	ExtSystemRepo interfaces.IExtSystemRepository
}

func (service *ExtSystemService) Create(extSystem bo.ExtSystemBO) (string, error) {
	extSystemDAO := dao.ExtSystemDAO{}
	extSystemDAO.FromBO(extSystem)

	return service.ExtSystemRepo.InsertOne(extSystemDAO)
}

func (service *ExtSystemService) Exist(extSystemID string) (bool, error) {
	return service.ExtSystemRepo.Exist(extSystemID)
}

func (service *ExtSystemService) List() ([]bo.ExtSystemBO, error) {
	list, err := service.ExtSystemRepo.SelectList()
	if err != nil {
		return nil, err
	}

	listBO := make([]bo.ExtSystemBO, 0, len(list))
	for _, item := range list {
		listBO = append(listBO, item.ToBO())
	}
	return listBO, nil
}
