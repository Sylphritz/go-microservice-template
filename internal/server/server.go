package server

import (
	"fmt"
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

func (s *Server) Start(port string) error {
	return http.ListenAndServe(fmt.Sprintf("localhost:%s", port), s.router)
}
