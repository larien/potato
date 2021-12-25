package main

import (
	"log"
	"net/http"

	"github.com/larien/go-skelethon/handler"
)

const (
	port = ":8080"
)

func main() {
	log.Println("starting server in port", port)

	http.HandleFunc("/potato", handler.V1GetPotato)

	http.ListenAndServe(port, nil)
}
