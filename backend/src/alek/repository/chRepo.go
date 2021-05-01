package repository

import (
	"alek/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type ChRepo struct{}

func NewChRepo() *ChRepo {
	return &ChRepo{}
}

var (
	ctx            = context.Background()
	sa             = option.WithCredentialsFile("../../../cultural-heritage-c8349-firebase-adminsdk.json")
	projectId      = "cultural-heritage-c8349"
	collectionName = "culturalHeritages"
)

func (*ChRepo) Save(ch *model.Ch) (*model.Ch, error) {
	client, err := firestore.NewClient(ctx, projectId, sa)
	if err != nil {
		log.Fatalf("Cannot connect with firestore: %v", err)
		return nil, err
	}

	_, _, err = client.Collection(collectionName).Add(ctx, ch)

	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}

	defer client.Close()
	return ch, nil
}
