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
	chrepo = repository.NewChRepo()
)

func (*ChService) Save(ch *model.Ch) (*model.Ch, error) {
	return chrepo.Save(ch)
}

func (*ChService) GetAll() ([]model.Ch, error) {
	return chrepo.GetAll()
}

func (*ChService) Search(search *model.Search) ([]model.Ch, error) {
	return chrepo.Search(search)
}

func (*ChService) Like(chId string) (*model.Ch, error) {
	return chrepo.Like(chId)
}

func (*ChService) Disike(chId string) (*model.Ch, error) {
	return chrepo.Dislike(chId)
}
