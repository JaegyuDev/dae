package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/3AM-Developer/dae/handlers"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "dae", log.LstdFlags)
	startDevHandler := handlers.NewStartDev(logger)

	r := mux.NewRouter() // Using Gorilla's Mux router here

	r.HandleFunc("/start-dev", startDevHandler.ServeHTTP) // Note the ServeHTTP method is being directly used

	server := &http.Server{
		Addr:         ":9090",
		Handler:      r, // Note that we are setting our Gorilla Mux router as the handler here
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)            // Buffered to prevent missing signals
	signal.Notify(sigChan, os.Interrupt, os.Kill) // Combined both signals into a single call

	sig := <-sigChan
	logger.Println("Received terminate, shutting down", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
