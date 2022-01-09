package router

import "net/http"

type Routes []Route

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
	IsAdmin bool
}
