package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Protected route (%v) is being accessed by %v", r.RequestURI, r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
