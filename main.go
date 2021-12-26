package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/larien/potato/config"
	"github.com/larien/potato/handler"
	"github.com/larien/potato/utils/router"
)

func main() {
	c := config.New()

	r := router.New(handler.Routes)

	log.Println("starting server in port", c.Server.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", c.Server.Port), r); err != nil {
		log.Fatal(err)
	}
}
