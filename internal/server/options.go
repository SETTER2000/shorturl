package server

import (
	"time"
)

// Option -.
type Option func(*Server)

// Port -.
//func Port(port string) Option {
//	return func(s *Server) {
//		s.server.Addr = net.JoinHostPort("", port)
//	}
//}

// Host -.
func Host(host string) Option {
	return func(s *Server) {
		s.server.Addr = host
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
