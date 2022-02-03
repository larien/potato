package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/larien/potato/internal/router/middlewares"
)

const adminPrefix = "/admin"

func New(routes Routes) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/health", healthcheck).Methods(http.MethodGet)

	for _, route := range routes {
		if route.IsAdmin {
			route.Handler = middlewares.Use(route.Handler, middlewares.IsAdmin)
			route.Path = adminPrefix + route.Path
		}

		route.Handler = middlewares.Use(route.Handler, route.Middlewares...)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
