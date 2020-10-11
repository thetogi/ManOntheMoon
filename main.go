package main

import (
	"ManOnTheMoonReviewService/route"
	"log"
	"net/http"
)

func main() {
	//TODO Safely start and stop server

	//TODO Make sure comments are formatted per Go Doc

	//TODO Add JWT token authetication
	//routes
	mux := route.GetRoutes()

	log.Printf("Starting Man on the Moon Game Rating Service on Port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != http.ErrServerClosed {
		// unexpected error. port in use?
		log.Fatalf("ListenAndServe(): %v", err)
	}
}
