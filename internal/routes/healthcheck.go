package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerHealthCheck(router *chi.Mux) {
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Okay"))
	})
}
