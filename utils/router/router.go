package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

const adminPrefix = "/admin"

func New(routes Routes) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/health", healthcheck).Methods(http.MethodGet)

	for _, route := range routes {
		if route.IsAdmin {
			route.Path = adminPrefix + route.Path
		}

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
