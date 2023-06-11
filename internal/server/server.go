// Package server реализует HTTP-сервер.
package server

import (
	"context"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/controller/grpc"
	"github.com/SETTER2000/shorturl/internal/controller/grpc/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 50 * time.Second
	_defaultAddr            = ":80"
	_defaultShutdownTimeout = 3 * time.Second
	_defaultHTTPS           = false
	_defaultCertFile        = ""
	_defaultKeyFile         = ""
	_defaultGrpcPort        = ""
)

// Server -.
type Server struct {
	certFile        string
	notify          chan error
	keyFile         string
	grpcPort        string
	isHTTPS         bool
	server          *http.Server
	srv             *grpc.Server
	cfg             *config.Config
	shutdownTimeout time.Duration
}

// New -.
func New(handlerHTTP http.Handler, opts ...Option) *Server {
	var logger = logrus.New()
	httpServer := &http.Server{
		Handler:      handlerHTTP,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	grpcServer := grpc.NewServer(grpc.Deps{
		Logger:  logger,
		Handler: &handler.IShorturlServer{},
	})
	s := &Server{
		server:          httpServer,
		srv:             grpcServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		isHTTPS:         _defaultHTTPS,
		certFile:        _defaultCertFile,
		keyFile:         _defaultKeyFile,
		grpcPort:        _defaultGrpcPort,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	switch s.isHTTPS {
	case true:
		go func() {
			s.notify <- s.server.ListenAndServeTLS(s.certFile, s.keyFile)
			close(s.notify)
		}()
	case false:
		go func() {
			s.notify <- s.server.ListenAndServe()
			close(s.notify)
		}()
	}

	go func() {
		s.srv.Logger.Info("Starting gRPC Server, PORT :", s.grpcPort)
		s.notify <- s.srv.ListenAndServer(s.grpcPort)
		close(s.notify)
	}()

}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
