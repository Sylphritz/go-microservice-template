package v1

import (
	"fmt"

	"github.com/go-chi/chi/v5"
)

const ApiVersion = "v1"

func RegisterRoutes(r chi.Router) {
	router := chi.NewRouter()

	registerPublicRoutes(router)
	registerProtectedRoutes(router)

	r.Mount(fmt.Sprintf("/%s", ApiVersion), router)
}
