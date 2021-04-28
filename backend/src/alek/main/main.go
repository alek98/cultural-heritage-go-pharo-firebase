package main

import "fmt"

func main() {
	fmt.Println("Starting Backend")
	/*
		1. cd src/alek/main
		2. open powershell (not stupid cmd)
		3. run `firebase emulators:start`
		3. run `$env:FIRESTORE_EMULATOR_HOST="localhost:8080"`
		4. run `go run .`
		5. start coding
	*/
	// startServer()

	testFirebase()
}
