package middlewares

import "net/http"

func Use(handler http.HandlerFunc, middlewares ...func(next http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
