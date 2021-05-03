package controller

import (
	"alek/model"
	"alek/service"
	"encoding/json"
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

func (*ChController) GetAll(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	chs, err := myservice.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chs)
}

func (*ChController) Search(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var search model.Search
	err := json.NewDecoder(request.Body).Decode(&search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling object")
		return
	}

	chs, err := myservice.Search(&search)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chs)
}

func (*ChController) Like(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ch model.Ch

	err := json.NewDecoder(request.Body).Decode(&ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling object")
		return
	}

	result, err := myservice.Like(&ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*ChController) Disike(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ch model.Ch

	err := json.NewDecoder(request.Body).Decode(&ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling object")
		return
	}

	result, err := myservice.Disike(&ch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
