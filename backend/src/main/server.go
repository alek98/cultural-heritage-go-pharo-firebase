package main

import (
	"controller"
	"fmt"
	"net/http"
	"router"
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
