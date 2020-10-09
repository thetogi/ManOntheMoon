package main

import (
	"ManOnTheMoonReviewService/route"
	"ManOnTheMoonReviewService/util"
	"log"
	"net/http"
)

func main() {
	//TODO Safely start and stop server

	//TODO Set Environment Variables
	util.Init()

	//TODO Set middlewares through alice

	//routes
	mux := route.GetRoutes()

	log.Printf("Staring Basic Web Server on Port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != http.ErrServerClosed {
		// unexpected error. port in use?
		log.Fatalf("ListenAndServe(): %v", err)
	}
}
