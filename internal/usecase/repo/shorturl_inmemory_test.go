package repo

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
)

type wantErr interface{}

func TestInMemory_Put(t *testing.T) {
	type fields struct {
		cfg  *config.Config
		m    map[string]entity.Shorturls
		lock *sync.Mutex
	}
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}
	tests := []struct {
		fields     fields
		name       string
		args       args
		UID        entity.UserID
		url        entity.URL
		countShort int
		countUID   int
		del        bool
	}{
		{
			name: "positive test #1",
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{Slug: "qwerty1", UserID: "uid-1", URL: "https://example1.com", Del: false},
			},
			countUID:   1,
			countShort: 1,
			UID:        "uid-1",
			url:        "https://example1.com",
			del:        false,
		}, {
			name: "positive test #2",
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{Slug: "qwerty2", UserID: "uid-2", URL: "https://example2.com", Del: false},
			},
			countUID:   2,
			countShort: 1,
			UID:        "uid-2",
			url:        "https://example2.com",
			del:        false,
		}, {
			name: "positive test #3",
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{Slug: "qwerty3", UserID: "uid-2", URL: "https://example3.com", Del: false},
			},
			countUID:   2,
			countShort: 2,
			UID:        "uid-2",
			url:        "https://example3.com",
			del:        false,
		}, {
			name: "positive test #4",
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{Slug: "qwerty4", UserID: "uid-2", URL: "https://example4.com", Del: false},
			},
			countUID:   2,
			countShort: 3,
			UID:        "uid-2",
			url:        "https://example4.com",
			del:        false,
		}, {
			name: "update Del and URL #5",
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{Slug: "qwerty4", UserID: "uid-2", URL: "https://example5.com", Del: true},
			},
			countUID:   2,
			countShort: 3,
			UID:        "uid-2",
			url:        "https://example5.com",
			del:        true,
		},
	}
	s := &InMemory{
		m: make(map[entity.UserID]entity.Shorturls),
	}
	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Put(tt.args.ctx, tt.args.sh); err != nil {
				t.Errorf("Put() error = %v, countUID %v", err, nil)
			}
			cUID := len(s.m)
			if cUID != tt.countUID {
				t.Errorf("Len memory = %v, countUID %v\n", cUID, tt.countUID)
			}

			cSh := len(s.m[tt.UID])
			if cSh != tt.countShort {
				t.Errorf("Len shorts; expected: %v, actual: %v\n", tt.countShort, cSh)
				fmt.Printf("List Shorturls: %v\n", s.m[tt.UID])
			}
			shorts := s.m[tt.UID]
			if idx == 4 && tt.url != shorts[2].URL {
				t.Errorf("Short URL; expected: %v, actual: %v\n", tt.url, shorts[2].URL)
			}
			if idx == 4 && tt.del != shorts[2].Del {
				t.Errorf("Short Del; expected: %v, actual: %v\n", tt.del, shorts[2].Del)
			}
		})
	}
}

func TestInMemory_Delete(t *testing.T) {
	user := entity.User{
		UserID: "1674872720465761244B_5",
	}
	lst := entity.List{
		ShortURL: "https://example.com/1674872720465761244B_5",
		URL:      "https://example.com/go/to/home.html",
	}
	user.DelLink = []entity.Slug{
		"1674872720465761244B_5",
	}
	user.Urls = append(user.Urls, lst)

	type fields struct {
		cfg  *config.Config
		m    map[string]entity.Shorturls
		lock *sync.Mutex
	}
	type args struct {
		ctx context.Context
		u   *entity.User
	}
	tests := []struct {
		fields     fields
		name       string
		args       args
		UID        entity.UserID
		url        entity.URL
		countShort int
		countUID   int
		del        bool
	}{
		{
			name: "positive test #4",
			args: args{
				ctx: context.Background(),
				u:   &user,
			},
			countUID:   2,
			countShort: 3,
			UID:        "uid-2",
			url:        "https://example4.com",
			del:        false,
		}, {
			name: "update Del and URL #5",
			args: args{
				ctx: context.Background(),
				u:   &user,
			},
			countUID:   2,
			countShort: 3,
			UID:        "uid-2",
			url:        "https://example5.com",
			del:        true,
		},
	}
	s := &InMemory{
		m: make(map[entity.UserID]entity.Shorturls),
	}
	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.Delete(tt.args.ctx, tt.args.u); err != nil {
				t.Errorf("Delete() error = %v, countUID %v", err, nil)
			}
			shorts := s.m[tt.UID]
			if idx == 4 && tt.url != shorts[2].URL {
				t.Errorf("Short URL; expected: %v, actual: %v\n", tt.url, shorts[2].URL)
			}
			if idx == 4 && tt.del != shorts[2].Del {
				t.Errorf("Short Del; expected: %v, actual: %v\n", tt.del, shorts[2].Del)
			}
		})
	}
}

func TestInMemory_Read(t *testing.T) {
	type fields struct {
		cfg  *config.Config
		m    map[entity.UserID]entity.Shorturls
		lock *sync.Mutex
	}
	tests := []struct {
		fields  fields
		wantErr wantErr
		name    string
	}{
		{
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				m:   tt.fields.m,
				cfg: tt.fields.cfg,
			}
			if err := s.Read(); (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemory_Get(t *testing.T) {
	sh := entity.Shorturl{
		Slug: "slug2", UserID: "uid2", URL: "https://example2.com", Del: false,
	}
	sh2 := entity.Shorturl{
		Slug: "slug3", UserID: "uid3", URL: "https://example3.com", Del: false,
	}
	type fields struct {
		lock *sync.Mutex
		m    map[entity.UserID]entity.Shorturls
		cfg  *config.Config
	}
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}

	fld := fields{
		m: make(map[entity.UserID]entity.Shorturls),
	}
	fld.m[sh.UserID] = append(fld.m[sh.UserID], sh)

	tests := []struct {
		want    *entity.Shorturl
		fields  fields
		args    args
		name    string
		wantErr bool
	}{
		{
			name:   "positive test #1",
			fields: fld,
			args: args{
				ctx: context.Background(),
				sh:  &sh,
			},
			want:    &sh,
			wantErr: false,
		}, {
			name:   "negative test #2",
			fields: fld,
			args: args{
				ctx: context.Background(),
				sh:  &sh2,
			},
			want:    &sh2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				m:   tt.fields.m,
				cfg: tt.fields.cfg,
			}
			got, err := s.Get(tt.args.ctx, tt.args.sh)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Printf("Mem slice: %v\n", s.m)
			fmt.Printf("Return func : %v err: %v\n", got, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemory_Post(t *testing.T) {

	type fields struct {
		cfg  *config.Config
		m    map[entity.UserID]entity.Shorturls
		lock *sync.Mutex
	}
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}
	tests := []struct {
		fields  fields
		args    args
		name    string
		wantErr bool
	}{
		{
			name: "positive test #1",
			fields: fields{
				m: make(map[entity.UserID]entity.Shorturls),
			},
			args: args{
				ctx: context.Background(),
				sh:  &entity.Shorturl{Slug: "qwerty2", UserID: "uid-2", URL: "https://example2.com", Del: false},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				m:   tt.fields.m,
				cfg: tt.fields.cfg,
			}
			if err := s.Post(tt.args.ctx, tt.args.sh); (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemory_Save(t *testing.T) {
	type fields struct {
		cfg  *config.Config
		m    map[entity.UserID]entity.Shorturls
		lock *sync.Mutex
	}
	tests := []struct {
		fields  fields
		name    string
		wantErr bool
	}{
		{
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				m:   tt.fields.m,
				cfg: tt.fields.cfg,
			}
			if err := s.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewInMemory(t *testing.T) {
	inMem := &InMemory{
		m: make(map[entity.UserID]entity.Shorturls),
	}
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		args args
		want *InMemory
		name string
	}{
		{
			name: "positive test #1",
			want: inMem,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemory(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
