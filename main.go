package main

import (
	"ManOnTheMoonReviewService/route"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//Add API Endpoints/Routes to be handled
	mux := route.SetRoutes()

	server := &http.Server{
		Handler:      mux,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Man on the Moon Game Rating Service on Port 8080")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Shutdown Server gracefully
	waitForShutdown(server)
}

//Wait for shutdown handles shutting down the server to close out connections when the os terminates the service
func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
