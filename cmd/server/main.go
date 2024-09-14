package main

import (
	"log"

	"github.com/sylphritz/go-microservice-template/internal/server"
)

func main() {
	s := server.New()

	if err := s.Start("4280"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
