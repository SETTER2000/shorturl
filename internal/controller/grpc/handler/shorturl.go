package handler

import (
	"context"
	"fmt"
	pb "github.com/SETTER2000/shorturl-service-api/api"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IShorturlServer -.
type IShorturlServer struct {
	pb.UnimplementedIShorturlServer
	service usecase.IShorturlRepo
}

// NewIShorturlHandler -.
func NewIShorturlHandler(h usecase.IShorturlRepo) *IShorturlServer {
	return &IShorturlServer{service: h}
}

// Post -.
func (s *IShorturlServer) Post(ctx context.Context, in *pb.PostRequest) (*pb.PostResponse, error) {
	var response pb.PostResponse
	logrus.Infof("IN входящие данные: %+v\n", in)
	sh := entity.Shorturl{
		URL: entity.URL(in.Shorturl.Url),
	}
	err := s.service.Post(ctx, &sh)
	if err != nil {
		response.Error = fmt.Sprintf("Пользователь со Slug %s уже существует, ну типа того", in.Shorturl.Slug)
	}

	//logrus.Infof("Ошибок нет. Ответ: %+v\n", res)

	return &response, nil
}

// LongLink -.
func (s *IShorturlServer) LongLink(ctx context.Context, in *pb.LongLinkRequest) (*pb.LongLinkResponse, error) {
	var response pb.LongLinkResponse
	logrus.Infof("IN входящие данные: %+v\n", in)
	sh := entity.Shorturl{
		URL:  entity.URL(in.Shorturl.Url),
		Slug: entity.Slug(in.Shorturl.Slug),
		//UserID: entity.UserID(entity.UserID(ctx.Value(s.cfg.AccessTokenName).(string))),
		Del: false,
	}
	err := s.service.Put(ctx, &sh)
	if err != nil {
		response.Error = fmt.Sprintf("ERRRRRRR::: %s ", err)
	}

	//logrus.Infof("Ошибок нет. Ответ: %+v\n", res)

	response.Shorturl = string(sh.URL)
	return &response, nil
}

// ShortLink -.
func (s *IShorturlServer) ShortLink(ctx context.Context, in *pb.ShortLinkRequest) (*pb.ShortLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortLink not implemented")
}

// UserAllLink -.
func (s *IShorturlServer) UserAllLink(ctx context.Context, in *pb.UserAllLinkRequest) (*pb.UserAllLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAllLink not implemented")
}

// AllLink -.
func (s *IShorturlServer) AllLink(ctx context.Context, in *pb.AllLinkRequest) (*pb.AllLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllLink not implemented")
}

// AllUsers -.
func (s *IShorturlServer) AllUsers(ctx context.Context, in *pb.AllUsersRequest) (*pb.AllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllUsers not implemented")
}

// UserDelLink -.
func (s *IShorturlServer) UserDelLink(ctx context.Context, in *pb.UserDelLinkRequest) (*pb.UserDelLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelLink not implemented")
}

// ReadService -.
func (s *IShorturlServer) ReadService(ctx context.Context, in *pb.ReadServiceRequest) (*pb.ReadServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadService not implemented")
}

// SaveService -.
func (s *IShorturlServer) SaveService(ctx context.Context, in *pb.SaveServiceRequest) (*pb.SaveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveService not implemented")
}
