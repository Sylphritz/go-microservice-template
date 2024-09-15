package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sylphritz/go-microservice-template/internal/config"
	"github.com/sylphritz/go-microservice-template/internal/server"
)

func main() {
	// Load configuration
	config.Load()

	// Create a server instance
	s := server.New(config.C.ServerHost, config.C.ServerPort)

	// Start the server in a goroutine to prevent the channel from blocking it
	go func() {
		log.Printf("Starting %v...", config.C.ServiceName)
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Create a channel of type `os.Signal` to listen for OS signals
	quit := make(chan os.Signal, 1)

	// Push incoming specified OS signals to the `quit` channel
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block the execution of the next lines until the `quit` channel is able to output a signal
	<-quit

	log.Printf("%v is shutting down...\n", config.C.ServiceName)

	// TODO: Close database connection and other resources like queues or files
	// ...

	// Create a background context that waits for up to 5 seconds for cleanup
	// If 5 seconds have passed and the server still hasn't shutdown, it will be forced to shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server has been forced to shutdown due to an error: %v", err)
	}

	log.Println("Server has been gracefully shut down")
}
