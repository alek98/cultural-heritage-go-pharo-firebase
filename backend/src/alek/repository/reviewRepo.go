package repository

import (
	"alek/model"
	"log"

	"cloud.google.com/go/firestore"
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

func (*ReviewRepo) RateReview(reviewId string, newRating int64) (*model.Review, error) {
	docRef := client.Collection(colReviews).Doc(reviewId)
	doc, _ := docRef.Get(ctx)
	var review model.Review
	doc.DataTo(&review)

	//calc ratings
	review.Rating = review.Rating*float64(review.TotalRatings) + float64(newRating)
	review.TotalRatings++
	review.Rating = review.Rating / float64(review.TotalRatings)

	_, err := docRef.Update(ctx, []firestore.Update{
		{
			Path:  "rating",
			Value: review.Rating,
		},
		{
			Path:  "totalRatings",
			Value: review.TotalRatings,
		},
	})

	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}

	return &review, nil
}
