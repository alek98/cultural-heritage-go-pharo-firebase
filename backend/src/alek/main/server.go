package main

import (
	"alek/controller"
	"alek/repository"
	"alek/router"
	"fmt"
	"net/http"
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
	myrouter.POST("/chs-search", chController.Search)
	myrouter.POST("/like", chController.Like)
	myrouter.POST("/dislike", chController.Disike)

	myrouter.SERVE(port)
}

func closeFirebaseConnection() {
	client := repository.GetClient()
	client.Close()
}
