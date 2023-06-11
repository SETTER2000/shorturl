package grpc

import (
	"fmt"
	pb "github.com/SETTER2000/shorturl-service-api/api"
	"github.com/SETTER2000/shorturl/internal/controller/grpc/handler"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

// Deps -.
type Deps struct {
	Logger *logrus.Logger

	Handler *handler.IShorturlServer
}

// Server -.
type Server struct {
	Deps
	srv *grpc.Server
}

// NewServer -.
func NewServer(deps Deps) *Server {
	logrus := logrus.NewEntry(deps.Logger)
	return &Server{
		srv: grpc.NewServer(
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_logrus.StreamServerInterceptor(logrus),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_logrus.UnaryServerInterceptor(logrus),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		),
		Deps: deps,
	}
}

// ListenAndServer -.
func (s *Server) ListenAndServer(port string) error {
	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	pb.RegisterIShorturlServer(s.srv, s.Deps.Handler)

	if err := s.srv.Serve(lis); err != nil {
		return err
	}

	return nil
}

// Stop -.
func (s *Server) Stop() {
	s.srv.GracefulStop()
}
