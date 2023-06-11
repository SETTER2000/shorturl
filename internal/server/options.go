package server

import (
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// Option -.
type Option func(*Server)

// Host -.
func Host(host string) Option {
	log.Printf("HOST::: %s", host)
	return func(s *Server) {
		domain := strings.Split(host, ":")
		s.server.Addr = net.JoinHostPort(domain[0], domain[1])
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

// PortGRPC -.
func PortGRPC(port string) Option {
	return func(s *Server) {
		s.grpcPort = port
	}
}

//
//// EnableGRPC - включить поддержку gRPC.
//func EnableGRPC(cfg *config.GRPC) Option {
//
//	grpcSrv := grpc.NewServer(grpc.Deps{
//		EntityHandler:
//	})
//
//	return func(s *Server) {
//		s.grpcSrv = grpcSrv
//	}
//}

// EnableHTTPS - опция подключает возможность использования SSL/TLS на сервере.
func EnableHTTPS(cfg *config.HTTP) Option {
	log.Printf("cfg.ServerDomain: %s\n", cfg.ServerDomain)
	// конструируем менеджер TLS-сертификатов
	manager := &autocert.Manager{
		// директория для хранения сертификатов
		Cache: autocert.DirCache("cache_dir"),
		// функция, принимающая Terms of Service издателя сертификатов
		Prompt: autocert.AcceptTOS,
		// перечень доменов, для которых будут поддерживаться сертификаты
		HostPolicy: autocert.HostWhitelist(cfg.ServerDomain, fmt.Sprintf("www.%s", cfg.ServerDomain)),
	}
	return func(s *Server) {
		flag := cfg.EnableHTTPS
		_, ok := os.LookupEnv("ENABLE_HTTPS")

		if !ok {
			if flag {
				s.isHTTPS = flag
				s.server.Addr = ":443"
				s.server.TLSConfig = manager.TLSConfig()
				s.certFile = fmt.Sprintf("%s/%s", cfg.CertsDir, cfg.CertFile)
				s.keyFile = fmt.Sprintf("%s/%s", cfg.CertsDir, cfg.KeyFile)
				log.Printf("enabled HTTPS: %v\n", flag)
			} else {
				s.isHTTPS = flag
				s.certFile = ""
				s.keyFile = ""
				log.Printf("disabled HTTPS:::%v\n", flag)
			}
		} else {
			s.isHTTPS = flag
			log.Printf("connect HTTPS: %v\n", flag)
		}
	}
}
