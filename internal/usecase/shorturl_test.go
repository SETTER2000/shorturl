package usecase

import (
	"context"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	shorturlRepo := mocks.NewIShorturlRepo(t)
	type args struct {
		r IShorturlRepo
	}
	tests := []struct {
		name string
		args args
		want *ShorturlUseCase
	}{
		{
			name: "checking ShorturlUsecase Layer creation, test #1",
			args: args{
				r: shorturlRepo,
			},
			want: New(shorturlRepo),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShorturlUseCase_Post(t *testing.T) {
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}

	sh := &entity.Shorturl{
		Slug:   "s1",
		URL:    "http://xx.ru",
		UserID: "uid1",
		Del:    false,
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				sh:  sh,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)
			shorturlRepo.
				On("Post", mock.Anything, tt.args.sh).
				Once().         // выполняется один раз
				Return(tt.want) // здесь конкретно указываем, что должен вернуть мок,
			// после того как его вызвали

			uc := &ShorturlUseCase{
				repo: shorturlRepo,
			}

			err := uc.Post(tt.args.ctx, tt.args.sh)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestShorturlUseCase_LongLink(t *testing.T) {
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}

	sh := &entity.Shorturl{
		Slug:   "s1",
		URL:    "http://xx.ru",
		UserID: "uid1",
		Del:    false,
	}

	sh2 := &entity.Shorturl{
		Slug:   "",
		URL:    "http://xx2.ru",
		UserID: "",
		Del:    false,
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				sh:  sh,
			},
			want:    sh.Slug,
			wantErr: false,
		},
		{
			name: "negative test #1",
			args: args{
				ctx: context.Background(),
				sh:  sh2,
			},
			want:    sh2.Slug,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)
			shorturlRepo.
				On("Put", mock.Anything, tt.args.sh).
				Once().     // выполняется один раз
				Return(nil) // здесь указываем, что должен вернуть мок, после того как его вызвали

			uc := &ShorturlUseCase{
				repo: shorturlRepo,
			}

			got, err := uc.LongLink(tt.args.ctx, tt.args.sh)
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
	shorturlRepo := mocks.NewIShorturlRepo(t)
	type fields struct {
		repo IShorturlRepo
	}
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Shorturl
		wantErr interface{}
	}{
		{
			name:   "positive test ShortLink #1",
			fields: fields{repo: shorturlRepo},
			args: args{
				ctx: context.Background(),
				sh: &entity.Shorturl{
					Slug:   "1",
					URL:    "http://xxzz",
					UserID: "1",
					Del:    false,
				}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo.
				On("Get", mock.Anything, tt.args.sh).
				Once().                     // выполняется один раз
				Return(tt.want, tt.wantErr) // здесь конкретно указываем, что должен вернуть мок,
			// после того как его вызвали

			uc := &ShorturlUseCase{
				repo: tt.fields.repo,
			}

			got, err := uc.ShortLink(tt.args.ctx, tt.args.sh)
			if err != tt.wantErr {
				t.Errorf("ShortLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestShorturlUseCase_ShortLinkError(t *testing.T) {
	shorturlRepo := mocks.NewIShorturlRepo(t)
	type fields struct {
		repo IShorturlRepo
	}
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Shorturl
		wantErr interface{}
	}{
		{
			name:   "negative test ShortLink #1",
			fields: fields{repo: shorturlRepo},
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{}},
			wantErr: ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo.
				On("Get", mock.Anything, tt.args.sh).
				Once().                 // выполняется один раз
				Return(nil, tt.wantErr) // здесь конкретно указываем, что должен вернуть мок,
			// после того как его вызвали

			uc := &ShorturlUseCase{
				repo: tt.fields.repo,
			}

			got, err := uc.ShortLink(tt.args.ctx, tt.args.sh)
			if err != tt.wantErr {
				t.Errorf("ShortLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}
