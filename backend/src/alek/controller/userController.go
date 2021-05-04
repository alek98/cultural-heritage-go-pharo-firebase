package controller

import (
	"alek/service"
	"encoding/json"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

var (
	userService = service.NewUserService()
)

func (*UserController) RateUser(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj := map[string]interface{}{
		"userName":  nil,
		"newRating": 0,
	}

	err := json.NewDecoder(request.Body).Decode(&obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error unmarshalling object")
		return
	}

	user, err := userService.RateUser(obj["userName"].(string), obj["newRating"].(float64))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
