package main

import (
	"fmt"
	"github.com/SETTER2000/shorturl/internal/app/config"
	"github.com/SETTER2000/shorturl/internal/app/shorturl"
	"github.com/SETTER2000/shorturl/internal/app/shorturl/db"
	"github.com/go-chi/chi/v5"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	router := chi.NewRouter()
	repository := db.NewRepository()
	shorturlHandler := shorturl.NewHandler(repository)
	shorturlHandler.Register(router)
	cfg := config.GetConfig()

	// start server
	start(router, cfg)
}

func start(router *chi.Mux, cfg *config.Config) {
	log.Println("start application")

	var listener net.Listener
	var listenErr error

	log.Println("listen tcp")

	// 0.0.0.0 - все ip, а не только 127.0.0.1
	// 127.0.0.1 - это ip адрес loop back интерфейса, в linux|macOS это lo интерфейс
	// 127.0.0.1 - это не текущая машина и не localhost
	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s",
		cfg.Listen.BindIP, cfg.Listen.Port))
	log.Printf("server is listener port %s:%s\n", cfg.Listen.BindIP, cfg.Listen.Port)

	if listenErr != nil {
		log.Println(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second, // таймаут на запись
		ReadTimeout:  15 * time.Second, // таймаут на чтение
	}
	log.Fatal(server.Serve(listener))
}
