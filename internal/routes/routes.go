package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sylphritz/go-microservice-template/internal/middleware"
	v1 "github.com/sylphritz/go-microservice-template/internal/routes/v1"
)

const ApiRoutePrefix = "/api"

func RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// For pinging
	registerHealthCheck(router)

	// V1 API routes
	router.Route(ApiRoutePrefix, func(r chi.Router) {
		v1.RegisterRoutes(r)
	})

	return router
}
