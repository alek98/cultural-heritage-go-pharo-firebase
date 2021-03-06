package main

import (
	"alek/controller"
	"alek/repository"
	"alek/router"
	"fmt"
	"net/http"
)

func startServer() {
	port := ":8000"
	myrouter := router.NewMuxRouter()
	myrouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Working")
	})

	chController := controller.NewChController()
	reviewController := controller.NewReviewController()
	commentController := controller.NewCommentController()
	userController := controller.NewUserController()
	myrouter.POST("/chs", chController.Save)
	myrouter.GET("/chs", chController.GetAll)
	myrouter.POST("/chs-search", chController.Search)
	myrouter.POST("/like", chController.Like)
	myrouter.POST("/dislike", chController.Disike)
	myrouter.GET("/reviews", reviewController.GetAll)
	myrouter.POST("/reviews/rate", reviewController.RateReview)
	myrouter.GET("/comments", commentController.GetAll)
	myrouter.POST("/comments", commentController.Save)
	myrouter.POST("/users/rate", userController.RateUser)
	myrouter.GET("/users", userController.GetAll)

	myrouter.SERVE(port)
}

func closeFirebaseConnection() {
	client := repository.GetClient()
	client.Close()
}
