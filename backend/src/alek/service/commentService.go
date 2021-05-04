package service

import (
	"alek/model"
	"alek/repository"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

var (
	commentRepo = repository.NewCommentRepo()
)

func (*CommentService) GetAll() ([]model.Comment, error) {
	return commentRepo.GetAll()
}

func (*CommentService) Save(comment *model.Comment) (*model.Comment, error) {
	return commentRepo.Save(comment)
}
