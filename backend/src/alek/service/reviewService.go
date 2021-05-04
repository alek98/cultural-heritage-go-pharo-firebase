package service

import (
	"alek/model"
	"alek/repository"
	"errors"
)

type ReviewService struct{}

func NewReviewService() *ReviewService {
	return &ReviewService{}
}

var (
	reviewRepo = repository.NewReviewRepo()
)

func (*ReviewService) GetAll() ([]model.Review, error) {
	return reviewRepo.GetAll()
}

func (*ReviewService) RateReview(reviewId string, newRating float64) (*model.Review, error) {
	if newRating == 1 || newRating == 2 || newRating == 3 || newRating == 4 || newRating == 5 {
		return reviewRepo.RateReview(reviewId, newRating)

	} else {
		err := errors.New("rating can take values: 1,2,3,4,5")
		return nil, err
	}
}
