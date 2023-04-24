package repo

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
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
