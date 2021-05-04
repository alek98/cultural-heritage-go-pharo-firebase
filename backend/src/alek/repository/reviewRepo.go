package repository

import (
	"alek/model"

	"google.golang.org/api/iterator"
)

type ReviewRepo struct{}

func NewReviewRepo() *ReviewRepo {
	return &ReviewRepo{}
}

func (*ReviewRepo) GetAll() ([]model.Review, error) {
	var reviews []model.Review
	iter := client.Collection(colReviews).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var review model.Review
		doc.DataTo(&review)
		review.Id = doc.Ref.ID
		reviews = append(reviews, review)
	}
	return reviews, nil
}
