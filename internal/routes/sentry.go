package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterSentryPanicRoute(router *chi.Mux) {
	router.Get("/testpanic", func(w http.ResponseWriter, r *http.Request) {
		panic("lol we ded")
	})
}
