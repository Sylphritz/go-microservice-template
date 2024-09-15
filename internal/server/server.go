package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sylphritz/go-microservice-template/internal/routes"
)

type Server struct {
	httpServer *http.Server
}

func New(host string, port int) *Server {
	addr := fmt.Sprintf("%v:%v", host, port)
	srv := &http.Server{
		Addr:    addr,
		Handler: routes.RegisterRoutes(),
	}

	return &Server{srv}
}

func (s *Server) Start() error {
	log.Printf("Server successfully started: %v\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
