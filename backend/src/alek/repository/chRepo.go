package repository

import (
	"alek/model"
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ChRepo struct{}

func NewChRepo() *ChRepo {
	return &ChRepo{}
}

var (
	ctx = context.Background()
	// sa variable is obligatory if working with cloud firestore and not firestore emulator
	// sa             = option.WithCredentialsFile("../../../cultural-heritage-c8349-firebase-adminsdk.json")
	projectId      = "cultural-heritage-c8349"
	collectionName = "culturalHeritages"
)

func (*ChRepo) Save(ch *model.Ch) (*model.Ch, error) {

	// client, err := firestore.NewClient(ctx, projectId, sa)
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Cannot connect with firestore: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, ch)
	fmt.Println(ch)

	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}

	return ch, nil
}

func (*ChRepo) GetAll() ([]model.Ch, error) {
	client, _ := firestore.NewClient(ctx, projectId)
	defer client.Close()

	var chs []model.Ch
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var ch model.Ch
		doc.DataTo(&ch)
		ch.Id = doc.Ref.ID
		chs = append(chs, ch)
	}

	return chs, nil
}

//TODO: CREATE SEARCH FUNCTION
func (*ChRepo) Search(search *model.Search) ([]model.Ch, error) {
	client, _ := firestore.NewClient(ctx, projectId)
	defer client.Close()

	var chs []model.Ch
	collectionRef := client.Collection(collectionName)
	query := collectionRef.Where("name", "==", search.Name)
	query = query.Where("avgRating", ">=", search.AvgRatingFrom)
	query = query.Where("avgRating", "<=", search.AvgRatingTo)
	// query = query.Where("")

	iter := query.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var ch model.Ch
		doc.DataTo(&ch)
		ch.Id = doc.Ref.ID
		chs = append(chs, ch)
	}
	return chs, nil
}
