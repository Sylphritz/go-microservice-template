package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sylphritz/go-microservice-template/internal/middleware"
)

func TestProtectedRoute(router chi.Router) {
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blahhhsss"))
	})
}

func registerProtectedRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		TestProtectedRoute(r)
	})
}
