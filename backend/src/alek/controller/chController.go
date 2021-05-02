package controller

import (
	"alek/model"
	"alek/service"
	"encoding/json"
	"log"
	"net/http"
)

type ChController struct{}

func NewChController() *ChController {
	return &ChController{}
}

var (
	myservice = service.NewChService()
)

func (*ChController) Save(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var culturalHeritage model.Ch
	err := json.NewDecoder(request.Body).Decode(&culturalHeritage)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling object")
		log.Fatalf("Error unmarshalling object")
		return
	}

	result, err := myservice.Save(&culturalHeritage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
