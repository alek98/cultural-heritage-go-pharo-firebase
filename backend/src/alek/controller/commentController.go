package controller

import (
	"alek/model"
	"alek/service"
	"encoding/json"
	"net/http"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

var (
	commentService = service.NewCommentService()
)

func (*CommentController) GetAll(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments, err := commentService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error getting objects")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func (*CommentController) Save(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var comment model.Comment
	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling objects.")
		return
	}

	result, err := commentService.Save(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
