package main

import (
	"log"

	"github.com/sylphritz/go-microservice-template/internal/config"
	"github.com/sylphritz/go-microservice-template/internal/server"
)

func main() {
	config.Load()

	s := server.New()

	log.Printf("Starting %v...", config.C.ServiceName)
	if err := s.Start(config.C.ServerHost, config.C.ServerPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
