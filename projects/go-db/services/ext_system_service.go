package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/objects/bo"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
	"github.com/Alexplusm/bazaa/projects/go-db/repositories"
)

type ExtSystemService struct {
	ExtSystemRepo repositories.ExtSystemRepository
}

func (service *ExtSystemService) CreateExtSystem(extSystem bo.ExtSystemBO) error {
	extSystemDAO := dao.ExtSystemDAO{}
	extSystemDAO.FromBO(extSystem)

	return service.ExtSystemRepo.InsertExtSystem(extSystemDAO)
}
