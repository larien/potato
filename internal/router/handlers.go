package router

import "net/http"

func healthcheck(w http.ResponseWriter, r *http.Request) {
	// A health check can assess anything that a microservice needs, including:
	// - dependencies
	// - system properties
	// - database connections
	// - endpoint connections
	// - resource availability

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status": "OK"}`))
}
