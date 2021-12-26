package router

import "net/http"

type Routes []Route

type Route struct {
	Pattern string
	Version string
	Method  string
	Handler http.HandlerFunc
}
