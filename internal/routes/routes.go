package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/sylphritz/go-microservice-template/internal/config"
	"github.com/sylphritz/go-microservice-template/internal/middleware"
	"github.com/sylphritz/go-microservice-template/internal/monitoring"
	v1 "github.com/sylphritz/go-microservice-template/internal/routes/v1"
)

const ApiRoutePrefix = "/api"

func RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	if config.C.SentryUrl != "" {
		router.Use(monitoring.InitSentry().Handle)
	}

	// For pinging
	registerHealthCheck(router)

	// For testing if Sentry works. Comment it out when you're done checking.
	// RegisterSentryPanicRoute(router)

	// V1 API routes
	router.Route(ApiRoutePrefix, func(r chi.Router) {
		v1.RegisterRoutes(r)
	})

	return router
}
