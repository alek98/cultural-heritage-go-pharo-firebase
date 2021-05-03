package repository

import (
	"alek/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ChRepo struct{}

func NewChRepo() *ChRepo {
	return &ChRepo{}
}

var (
	ctx            = context.Background()
	projectId      = "cultural-heritage-c8349"
	collectionName = "culturalHeritages"
)

func (*ChRepo) Save(ch *model.Ch) (*model.Ch, error) {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Cannot connect with firestore: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, ch)

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

func (*ChRepo) Search(search *model.Search) ([]model.Ch, error) {
	client, _ := firestore.NewClient(ctx, projectId)
	defer client.Close()

	var chs []model.Ch
	collectionRef := client.Collection(collectionName)
	query := collectionRef.Query
	// search
	if search.AvgRatingFrom != 0 {
		query = query.Where("avgRating", ">=", search.AvgRatingFrom)
	}
	if search.AvgRatingTo != 0 {
		query = query.Where("avgRating", "<=", search.AvgRatingTo)
	}
	if search.ChTypeName != "" {
		query = query.Where("chtype.name", "==", search.ChTypeName)
	}
	if search.Street != "" {
		query = query.Where("location.street", "==", search.Street)
	}
	if search.City != "" {
		query = query.Where("location.city", "==", search.City)
	}
	if search.Country != "" {
		query = query.Where("location.country", "==", search.Street)
	}
	if search.Name != "" {
		query = query.Where("name", "==", search.Name)
	}

	// first sort by avgRating is obligatory if any range  filter (avgRating) is activated
	// https://firebase.google.com/docs/firestore/query-data/order-limit-data#limitations
	if search.AvgRatingTo != 0 || search.AvgRatingFrom != 0 {
		query = query.OrderBy("avgRating", firestore.Asc)
	}

	// sort
	if search.Sort.SortByName != "" {
		if search.Sort.SortByName == "desc" {
			query = query.OrderBy("name", firestore.Desc)
		} else {
			query = query.OrderBy("name", firestore.Asc)
		}
	} else if search.Sort.SortByRating != "" {
		if search.Sort.SortByRating == "desc" {
			query = query.OrderBy("rating", firestore.Desc)
		} else {
			query = query.OrderBy("rating", firestore.Asc)
		}
	} else if search.Sort.SortByChTypeName != "" {
		if search.Sort.SortByChTypeName == "desc" {
			query = query.OrderBy("chtype.name", firestore.Desc)
		} else {
			query = query.OrderBy("chtype.name", firestore.Asc)
		}
	}

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
