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
