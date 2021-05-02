package main

import (
	"alek/controller"
	"alek/router"
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
)

func startServer() {
	port := ":8080"
	myrouter := router.NewMuxRouter()
	myrouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Working")
	})

	chController := controller.NewChController()
	myrouter.POST("/chs", chController.Save)
	myrouter.GET("/chs", chController.GetAll)

	myrouter.SERVE(port)
}

func testFirebase() {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "cultural-heritage-c8349")
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
