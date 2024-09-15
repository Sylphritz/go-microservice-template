package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sylphritz/go-microservice-template/internal/routes"
)

type Server struct {
	router *chi.Mux
}

func New() *Server {
	return &Server{
		router: routes.RegisterRoutes(),
	}
}

func (s *Server) Start(host string, port int) error {
	addr := fmt.Sprintf("%v:%v", host, port)
	log.Printf("Server successfully started: %v\n", addr)
	return http.ListenAndServe(addr, s.router)
}
