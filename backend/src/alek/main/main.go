package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Backend")
	/*
			How to run firebase (in cmd)
			1. cd firebase
			2. run `firebase emulators:start --export-on-exit=testdata --import=testdata`

		  How to run golang
			-----------------
			1. cd src/alek/main
			2. set  FIRESTORE_EMULATOR_HOST=localhost:8081
			3. run `go build && main.exe`
			4. start coding
	*/
	// testFirebase()
	startServer()
}
