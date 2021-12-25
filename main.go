package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	log.Println("starting server in port", port)

	http.HandleFunc("/potato", v1GetPotato)

	http.ListenAndServe(port, nil)
}

func v1GetPotato(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		w.Write([]byte("{\"content\": \"potato\"}"))
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("{\"error\": \"method not allowed\"}"))
}
