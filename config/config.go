package config

import (
	"flag"
	"log"

	"github.com/larien/potato/internal/config"
)

type Config struct {
	Server Server
}

type Server struct {
	Port int
}

func New() Config {
	filename := flag.String("config", "config", "configuration file, default config/config.yml")
	flag.Parse()

	var c Config
	if err := config.Parse(*filename, &c); err != nil {
		log.Fatal(err)
	}
	return c
}
