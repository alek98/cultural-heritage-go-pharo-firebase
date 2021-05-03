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

func (*ChService) GetAll() ([]model.Ch, error) {
	return repo.GetAll()
}

func (*ChService) Search(search *model.Search) ([]model.Ch, error) {
	return repo.Search(search)
}

func (*ChService) Like(chId string) (*model.Ch, error) {
	return repo.Like(chId)
}

func (*ChService) Disike(chId string) (*model.Ch, error) {
	return repo.Dislike(chId)
}
