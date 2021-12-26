package main

import (
	"log"
	"net/http"

	"github.com/larien/potato/handler"
)

const (
	port = ":8080"
)

func main() {
	log.Println("starting server in port", port)

	http.HandleFunc("/potato", handler.V1GetPotato)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
