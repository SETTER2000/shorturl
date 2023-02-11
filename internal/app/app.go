package app

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SETTER2000/shorturl/config"
	v1 "github.com/SETTER2000/shorturl/internal/controller/http/v1"
	"github.com/SETTER2000/shorturl/internal/server"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/internal/usecase/repo"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/go-chi/chi/v5"
)

func Run(cfg *config.Config) {
	// logger
	l := logger.New(cfg.Log.Level)
	// seed
	rand.Seed(time.Now().UnixNano())

	file, _ := os.OpenFile(cfg.FileStorage, os.O_RDONLY|os.O_CREATE, 0777)

	// Use case
	shorturlUseCase := usecase.New(
		repo.NewProducer(file),
		repo.NewConsumer(file),
		repo.NewInMemory(),
	)

	// HTTP Server
	handler := chi.NewRouter()
	v1.NewRouter(handler, l, shorturlUseCase, cfg.HTTP)
	httpServer := server.New(handler, server.Host(cfg.HTTP.ServerAddress))

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
