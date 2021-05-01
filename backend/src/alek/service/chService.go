package service

import (
	"alek/model"
	"alek/repository"
)

type ChService struct{}

func NewChService() *ChService {
	return &ChService{}
}

var (
	repo = repository.NewChRepo()
)

func (*ChService) Save(ch *model.Ch) (*model.Ch, error) {
	return repo.Save(ch)
}
