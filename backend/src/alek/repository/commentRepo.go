package repository

import (
	"alek/model"
	"log"

	"google.golang.org/api/iterator"
)

type CommentRepo struct{}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{}
}

func (*CommentRepo) GetAll() ([]model.Comment, error) {
	var comments []model.Comment
	iter := client.Collection(colComments).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var comment model.Comment
		doc.DataTo(&comment)
		comment.Id = doc.Ref.ID
		comments = append(comments, comment)
	}
	return comments, nil
}

func (*CommentRepo) Save(comment *model.Comment) (*model.Comment, error) {
	_, _, err := client.Collection(colComments).Add(ctx, map[string]interface{}{
		"content":  comment.Content,
		"reviewId": comment.ReviewId,
		"userName": comment.UserName,
	})
	if err != nil {
		log.Fatalf("Failed saving to firestore: %v", err)
		return nil, err
	}
	return comment, nil
}
