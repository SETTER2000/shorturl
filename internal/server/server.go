// Package server реализует HTTP-сервер.
package server

import (
	"context"
	"github.com/SETTER2000/shorturl/config"
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
)

// Server -.
type Server struct {
	isHTTPS         bool
	notify          chan error
	certFile        string
	keyFile         string
	server          *http.Server
	cfg             *config.HTTP
	shutdownTimeout time.Duration
}

// New -.
func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
		isHTTPS:         _defaultHTTPS,
		certFile:        _defaultCertFile,
		keyFile:         _defaultKeyFile,
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
