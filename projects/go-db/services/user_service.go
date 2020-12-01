package services

import (
	"github.com/Alexplusm/bazaa/projects/go-db/interfaces"
	"github.com/Alexplusm/bazaa/projects/go-db/objects/dao"
)

type UserService struct {
	UserRepo interfaces.IUserRepository
}

func (service *UserService) CreateUser(userID string) error {
	userDAO := dao.UserDAO{UserID: userID}

	// TODO: why user already exist not error while insert?
	return service.UserRepo.InsertUser(userDAO)
}

func (service *UserService) UserExist(userID string) (bool, error) {
	return service.UserRepo.UserExist(userID)
}
