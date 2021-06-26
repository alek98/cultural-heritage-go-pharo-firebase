package service

import (
	"alek/model"
	"alek/repository"
	"errors"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

var (
	userRepo = repository.NewUserRepo()
)

func (*UserService) RateUser(userName string, newRating float64) (*model.User, error) {
	if newRating == 1 || newRating == 2 || newRating == 3 || newRating == 4 || newRating == 5 {
		return userRepo.RateUser(userName, newRating)

	} else {
		err := errors.New("rating can take values: 1,2,3,4,5")
		return nil, err
	}
}

func (*UserService) GetAll() ([]model.User, error) {
	return userRepo.GetAll()
}
