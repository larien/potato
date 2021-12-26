package router

import "net/http"

func New(routes Routes) {
	for _, route := range routes {
		http.HandleFunc(route.Pattern, route.Handler)
	}
}
