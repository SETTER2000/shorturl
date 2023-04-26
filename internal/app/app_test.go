package app_test

import (
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/app"
	"log"
	"os"
)

// Example using Chi router
func ExampleRun() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
		os.Exit(1)
	}

	app.Run(cfg)
}
