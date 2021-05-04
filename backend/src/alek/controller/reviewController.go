package controller

import (
	"alek/service"
	"encoding/json"
	"net/http"
)

type ReviewController struct{}

func NewReviewController() *ReviewController {
	return &ReviewController{}
}

var (
	reviewService = service.NewReviewService()
)

func (*ReviewController) GetAll(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reviews, err := reviewService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reviews)
}
