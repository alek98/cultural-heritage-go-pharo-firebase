package main

import (
	"alek/model"
	"alek/repository"
	"fmt"
)

func main() {
	fmt.Println("Starting Backend")
	/*
			How to run firebase (in cmd cannot run in powershell)
			1. cd firebase
			2. run `firebase emulators:start --export-on-exit=testdata --import=testdata`

		  How to run golang
			-----------------
			1. cd src/alek/main
			2. open powershell (not stupid cmd)
			3. run `$env:FIRESTORE_EMULATOR_HOST="localhost:8080"`
			4. run `go run .`
			5. start coding
	*/
	// testFirebase()
	// startServer()
	chrepo := repository.NewChRepo()
	chrepo.Save(&model.Ch{
		AvgRating: 4.5,
		ChType: model.ChType{
			Name:        "bilo koji tip",
			Description: "bilo koji description",
		},
		Description: "bilo sta",
		Location: model.Location{
			City:    "Novi Sad",
			Country: "Serbia",
			Street:  "neznana ulica",
		},
		Name: "Novi Tip",
	})
}
