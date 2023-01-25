package main

import (
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
