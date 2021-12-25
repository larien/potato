package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	log.Println("starting server in port", port)

	http.ListenAndServe(port, nil)
}
