// Package app - точка входа в проект, запуск сервиса shortener.
package app

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/SETTER2000/shorturl/scripts"

	"github.com/go-chi/chi/v5"
	"github.com/xlab/closer"

	"github.com/SETTER2000/shorturl/config"
	v1 "github.com/SETTER2000/shorturl/internal/controller/http/v1"
	"github.com/SETTER2000/shorturl/internal/server"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/internal/usecase/repo"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
)

// Run запуск сервиса
func Run(cfg *config.Config) {
	closer.Bind(cleanup)
	// logger
	l := logger.New(cfg.Log.Level)
	// seed
	rand.Seed(time.Now().UnixNano())

	// Use case
	var shorturlUseCase usecase.Shorturl
	if !scripts.CheckEnvironFlag("DATABASE_DSN", cfg.Storage.ConnectDB) {
		if cfg.FileStorage == "" {
			l.Warn("In memory storage!!!")
			shorturlUseCase = usecase.New(repo.NewInMemory(cfg))
		} else {
			l.Info("File storage - is work...")
			shorturlUseCase = usecase.New(repo.NewInFiles(cfg))
			if err := shorturlUseCase.ReadService(); err != nil {
				l.Error(fmt.Errorf("app - Read - shorturlUseCase.ReadService: %w", err))
			}
		}
	} else {
		// DB
		db, err := repo.New(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%e\n", err)
			os.Exit(1)
		}
		l.Info("DB SQL - is work...")
		shorturlUseCase = usecase.New(repo.NewInSQL(db))
	}

	// HTTP Server
	handler := chi.NewRouter()
	handler.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	v1.NewRouter(handler, l, shorturlUseCase, cfg)
	httpServer := server.New(handler, server.Host(cfg.HTTP.ServerAddress))

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		if err := shorturlUseCase.SaveService(); err != nil {
			l.Error(fmt.Errorf("app - Save - shorturlUseCase.SaveService: %w", err))
		}
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	closer.Hold()

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

func cleanup() {
	fmt.Print("Hang on! I'm closing some DBs, wiping some trails..")
	time.Sleep(1 * time.Second)
	fmt.Println("  Done...")
}
