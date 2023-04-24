package repo

import (
	"context"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"reflect"
	"sync"
	"testing"
)

func TestInMemory_Put(t *testing.T) {
	type fields struct {
		lock *sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	type args struct {
		ctx context.Context
		sh  *entity.Shorturl
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		countUID   int
		countShort int
		UID        string
		url        string
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
		m: make(map[string]entity.Shorturls),
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
		Slug: "1674872720465761244B_5",
		URL:  "https://example.com/go/to/home.html",
	}
	user.DelLink = []string{
		"1674872720465761244B_5",
	}
	user.Urls = append(user.Urls, lst)

	type fields struct {
		lock *sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	type args struct {
		ctx context.Context
		u   *entity.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		countUID   int
		countShort int
		UID        string
		url        string
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
		m: make(map[string]entity.Shorturls),
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

func TestInMemory_Get(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
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
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				m:    tt.fields.m,
				cfg:  tt.fields.cfg,
			}
			got, err := s.Get(tt.args.ctx, tt.args.sh)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemory_GetAll(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	type args struct {
		ctx context.Context
		u   *entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				m:    tt.fields.m,
				cfg:  tt.fields.cfg,
			}
			got, err := s.GetAll(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemory_Read(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				m:    tt.fields.m,
				cfg:  tt.fields.cfg,
			}
			if err := s.Read(); (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemory_Save(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				m:    tt.fields.m,
				cfg:  tt.fields.cfg,
			}
			if err := s.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemory_delete(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	type args struct {
		u *entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				m:    tt.fields.m,
				cfg:  tt.fields.cfg,
			}
			if err := s.delete(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemory_searchBySlug(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	type args struct {
		sh *entity.Shorturl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Shorturl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				m:    tt.fields.m,
				cfg:  tt.fields.cfg,
			}
			got, err := s.searchBySlug(tt.args.sh)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchBySlug() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBySlug() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemory_searchUID(t *testing.T) {
	type fields struct {
		lock sync.Mutex
		m    map[string]entity.Shorturls
		cfg  *config.Config
	}
	type args struct {
		sh *entity.Shorturl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Shorturl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &InMemory{
				lock: tt.fields.lock,
				m:    tt.fields.m,
				cfg:  tt.fields.cfg,
			}
			got, err := s.searchUID(tt.args.sh)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchUID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInMemory(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want *InMemory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInMemory(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
