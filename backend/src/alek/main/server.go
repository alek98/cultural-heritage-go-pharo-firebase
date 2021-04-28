package main

import (
	"alek/controller"
	"alek/router"
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func startServer() {
	port := ":8080"
	myrouter := router.NewMuxRouter()
	myrouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Working")
	})

	booksController := controller.NewBooksController()
	myrouter.GET("/books", booksController.GetAll)

	myrouter.SERVE(port)
}

func testFirebase() {

	sa := option.WithCredentialsFile("../../../cultural-heritage-c8349-firebase-adminsdk.json")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cultural-heritage-c8349", sa)
	if err != nil {
		log.Fatalf("ne radi %v", err)
	}
	defer client.Close()

	ch := client.Doc("culturalHeritages/FHu8NeAv5VsPR9Jjq00p")
	chSnapshot, err := ch.Get(ctx)
	if err != nil {
		fmt.Println("wtf has happend")
		return
	}
	fmt.Println(chSnapshot.Data())
}
