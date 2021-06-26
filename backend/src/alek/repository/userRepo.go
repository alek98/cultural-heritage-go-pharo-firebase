package repository

import (
	"alek/model"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}
func (*UserRepo) RateUser(userName string, newRating float64) (*model.User, error) {
	docRef := client.Collection(colUsers).Doc(userName)
	doc, _ := docRef.Get(ctx)
	var user model.User
	doc.DataTo(&user)

	//calc rating
	user.Rating = user.Rating*float64(user.TotalRatings) + newRating
	user.TotalRatings++
	user.Rating = user.Rating / float64(user.TotalRatings)

	_, err := docRef.Update(ctx, []firestore.Update{
		{
			Path:  "rating",
			Value: user.Rating,
		},
		{
			Path:  "totalRatings",
			Value: user.TotalRatings,
		},
	})

	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}

	return &user, nil
}

func (*UserRepo) GetAll() ([]model.User, error) {
	var users []model.User
	iter := client.Collection(colUsers).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var user model.User
		doc.DataTo(&user)
		user.Id = doc.Ref.ID
		users = append(users, user)
	}

	return users, nil
}
