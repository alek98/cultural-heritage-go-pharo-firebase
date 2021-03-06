package repository

import (
	"alek/model"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ChRepo struct{}

func NewChRepo() *ChRepo {
	return &ChRepo{}
}

func (*ChRepo) Save(ch *model.Ch) (*model.Ch, error) {
	_, _, err := client.Collection(colChs).Add(ctx, ch)

	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}

	return ch, nil
}

func (*ChRepo) GetAll() ([]model.Ch, error) {
	var chs []model.Ch
	iter := client.Collection(colChs).Documents(ctx)
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

	var chs []model.Ch
	collectionRef := client.Collection(colChs)
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
		query = query.Where("location.country", "==", search.Country)
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
	if search.SortByName != "" {
		if search.SortByName == "desc" {
			query = query.OrderBy("name", firestore.Desc)
		} else {
			query = query.OrderBy("name", firestore.Asc)
		}
	} else if search.SortByRating != "" {
		if search.SortByRating == "desc" {
			query = query.OrderBy("avgRating", firestore.Desc)
		} else {
			query = query.OrderBy("avgRating", firestore.Asc)
		}
	} else if search.SortByChTypeName != "" {
		if search.SortByChTypeName == "desc" {
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

func (*ChRepo) Like(chId string) (*model.Ch, error) {
	docRef := client.Collection(colChs).Doc(chId)
	doc, _ := docRef.Get(ctx)
	likes, _ := doc.DataAt("likes")
	_, err := docRef.Update(ctx, []firestore.Update{
		{
			Path:  "likes",
			Value: likes.(int64) + 1,
		},
	})

	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}

	var ch model.Ch
	doc, _ = docRef.Get(ctx)
	doc.DataTo(&ch)
	return &ch, nil
}

func (*ChRepo) Dislike(chId string) (*model.Ch, error) {
	docRef := client.Collection(colChs).Doc(chId)
	doc, _ := docRef.Get(ctx)
	dislikes, _ := doc.DataAt("dislikes")
	_, err := docRef.Update(ctx, []firestore.Update{
		{
			Path:  "dislikes",
			Value: dislikes.(int64) + 1,
		},
	})

	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}

	var ch model.Ch
	doc, _ = docRef.Get(ctx)
	doc.DataTo(&ch)
	return &ch, nil
}
