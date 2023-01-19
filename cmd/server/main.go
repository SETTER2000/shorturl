package main

import (
	"github.com/SETTER2000/shorturl/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Handlers)
	http.HandleFunc("/status", handlers.StatusHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
