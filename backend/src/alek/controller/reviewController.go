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

func (*ReviewController) RateReview(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj := map[string]interface{}{
		"reviewId":  nil,
		"newRating": 0,
	}
	err := json.NewDecoder(request.Body).Decode(&obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling object")
		return
	}

	review, err := reviewService.RateReview(obj["reviewId"].(string), obj["newRating"].(float64))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(review)
}
