package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/larien/potato/config"
	"github.com/larien/potato/handler"
)

func main() {
	c := config.New()

	log.Println("starting server in port", c.Server.Port)

	http.HandleFunc("/potato", handler.V1GetPotato)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", c.Server.Port), nil); err != nil {
		log.Fatal(err)
	}
}
