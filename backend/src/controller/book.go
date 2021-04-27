package controller

import (
	"encoding/json"
	"net/http"
)

type booksController struct{}

func NewBooksController() *booksController {
	return &booksController{}
}

func (*booksController) GetAll(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("knjige")
}
