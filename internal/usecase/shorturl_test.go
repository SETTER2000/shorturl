package usecase

import (
	"github.com/SETTER2000/shorturl/internal/entity"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		r ShorturlRepo
	}
	tests := []struct {
		name string
		args args
		want *ShorturlUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShorturlUseCase_LongLink(t *testing.T) {
	type fields struct {
		repo ShorturlRepo
	}
	type args struct {
		sh entity.Shorturl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &ShorturlUseCase{
				repo: tt.fields.repo,
			}
			got, err := uc.LongLink(tt.args.sh)
			if (err != nil) != tt.wantErr {
				t.Errorf("LongLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LongLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShorturlUseCase_ShortLink(t *testing.T) {
	type fields struct {
		repo ShorturlRepo
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &ShorturlUseCase{
				repo: tt.fields.repo,
			}
			got, err := uc.ShortLink(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ShortLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}
