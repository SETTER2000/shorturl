package v1

import (
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"net/http"
	"testing"
)

func Test_shorturlRoutes_longLink(t *testing.T) {
	type fields struct {
		s   usecase.Shorturl
		l   logger.Interface
		cfg config.HTTP
	}
	type args struct {
		res http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &shorturlRoutes{
				s:   tt.fields.s,
				l:   tt.fields.l,
				cfg: tt.fields.cfg,
			}
			r.longLink(tt.args.res, tt.args.req)
		})
	}
}

func Test_shorturlRoutes_shortLink(t *testing.T) {
	type fields struct {
		s   usecase.Shorturl
		l   logger.Interface
		cfg config.HTTP
	}
	type args struct {
		res http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &shorturlRoutes{
				s:   tt.fields.s,
				l:   tt.fields.l,
				cfg: tt.fields.cfg,
			}
			r.shortLink(tt.args.res, tt.args.req)
		})
	}
}
