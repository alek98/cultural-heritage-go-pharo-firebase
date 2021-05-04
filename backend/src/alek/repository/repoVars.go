package repository

import (
	"context"

	"cloud.google.com/go/firestore"
)

var (
	ctx       = context.Background()
	projectId = "cultural-heritage-c8349"
	client, _ = firestore.NewClient(ctx, projectId)

	// collection names
	colChs      = "culturalHeritages"
	colReviews  = "reviews"
	colComments = "comments"
)

func GetClient() *firestore.Client {
	return client
}
