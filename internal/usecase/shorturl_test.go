package usecase

import (
	"context"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

var cfg = &config.Config{
	HTTP: config.HTTP{
		ServerAddress: "localhost:8080",
		BaseURL:       "http://localhost:8080",
	},
	Storage: config.Storage{
		FileStorage: "storage.txt",
		ConnectDB:   "postgres://shorturl:DBshorten-2023@127.0.0.1:5432/shorturl?sslmode=disable",
	},
	Cookie: config.Cookie{
		AccessTokenName: "access_token",
		SecretKey:       "RtsynerpoGIYdab_s234r",
	},
}

func TestNew(t *testing.T) {
	shorturlRepo := mocks.NewIShorturlRepo(t)
	type args struct {
		r IShorturlRepo
	}
	tests := []struct {
		args args
		want *ShorturlUseCase
		name string
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
		args    args
		want    interface{}
		wantErr error
		name    string
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				sh:  sh,
			},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "negative test ShortLink #1",
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{}},
			wantErr: ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)
			shorturlRepo.
				On("Post", mock.Anything, tt.args.sh).
				Once().            // выполняется один раз
				Return(tt.wantErr) // указываем, что должен вернуть мок

			uc := &ShorturlUseCase{
				repo: shorturlRepo,
			}

			err := uc.Post(tt.args.ctx, tt.args.sh)
			if (err != nil) && (err != tt.wantErr) {
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
		URL:    "",
		UserID: "",
		Del:    false,
	}

	tests := []struct {
		args    args
		want    interface{}
		wantErr error
		name    string
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				sh:  sh,
			},
			want:    sh.Slug,
			wantErr: nil,
		},
		{
			name: "negative test #1",
			args: args{
				ctx: context.Background(),
				sh:  sh2,
			},
			want:    sh2.Slug,
			wantErr: ErrBadRequest,
		},
		{
			name: "negative test #2",
			args: args{
				ctx: context.Background(),
				sh:  sh2,
			},
			want:    "",
			wantErr: ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)

			shorturlRepo.
				On("Put", mock.Anything, tt.args.sh).
				Once().            // выполняется один раз
				Return(tt.wantErr) // здесь указываем, что должен вернуть мок, после того как его вызвали

			uc := &ShorturlUseCase{
				repo: shorturlRepo,
			}

			got, err := uc.LongLink(tt.args.ctx, tt.args.sh)
			if (err != nil) && (err != tt.wantErr) {
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
		fields  fields
		args    args
		want    *entity.Shorturl
		wantErr interface{}
		name    string
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
		fields  fields
		args    args
		want    *entity.Shorturl
		wantErr interface{}
		name    string
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

func TestShorturlUseCase_UserAllLink(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *entity.User
	}

	tests := []struct {
		args    args
		want    *entity.User
		wantErr error
		name    string
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{UserID: "uid_1"},
			},
			wantErr: nil,
		},
		{
			name: "negative test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{UserID: "uid_1"},
			},
			want:    nil,
			wantErr: ErrBadRequest,
		},
		{
			name: "negative test #2",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{},
			},
			want:    nil,
			wantErr: ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)
			shorturlRepo.
				On("GetAll", tt.args.ctx, tt.args.u).
				Times(1). // выполняется один раз
				Return(tt.want, tt.wantErr)

			uc := New(shorturlRepo)
			got, err := uc.UserAllLink(tt.args.ctx, tt.args.u)

			if (err != nil) && (err != tt.wantErr) {
				t.Errorf("UserAllLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserAllLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestShorturlUseCase_UserDelLink(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *entity.User
	}
	tests := []struct {
		args    args
		want    *entity.User
		wantErr error
		name    string
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{UserID: "1682704080950404852_x3"},
			},
			wantErr: nil,
		},
		{
			name: "negative test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{},
			},
			want:    nil,
			wantErr: ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)
			shorturlRepo.
				On("Delete", mock.Anything, tt.args.u).
				Once(). // выполняется один раз
				Return(tt.wantErr)

			uc := New(shorturlRepo)
			err := uc.UserDelLink(tt.args.ctx, tt.args.u)
			if (err != nil) && (err != tt.wantErr) {
				t.Errorf("UserDelLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("UserDelLink() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestShorturlUseCase_SaveService(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *entity.User
	}
	tests := []struct {
		args    args
		want    *entity.User
		wantErr error
		name    string
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{UserID: "1682704080950404852_x3"},
			},
		},
		{
			name: "negative test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{UserID: "uid_1"},
			},
			want:    nil,
			wantErr: ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)
			shorturlRepo.
				On("Save").
				Once(). // выполняется один раз
				Return(tt.wantErr)

			uc := New(shorturlRepo)
			err := uc.SaveService()
			if (err != nil) && (err != tt.wantErr) {
				t.Errorf("SaveService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("SaveService() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestShorturlUseCase_ReadService(t *testing.T) {
	type args struct {
		ctx context.Context
		u   *entity.User
	}
	tests := []struct {
		args    args
		want    *entity.User
		wantErr error
		name    string
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{UserID: "1682704080950404852_x3"},
			},
		},
		{
			name: "negative test #1",
			args: args{
				ctx: context.Background(),
				u:   &entity.User{UserID: "uid_1"},
			},
			want:    nil,
			wantErr: ErrBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shorturlRepo := mocks.NewIShorturlRepo(t)
			shorturlRepo.
				On("Read").
				Once(). // выполняется один раз
				Return(tt.wantErr)

			uc := New(shorturlRepo)
			err := uc.ReadService()
			if (err != nil) && (err != tt.wantErr) {
				t.Errorf("ReadService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("SaveService() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
