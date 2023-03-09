package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"sync"
)

// InMemory
// Если вам нужно защитить доступ к простой структуре данных, такой как слайс,
// или map, или что-нибудь своё, и если интерфейс доступа к этой структуре данных
// прост и прямолинеен — начинайте с мьютекса.
// Это также помогает спрятать «грязные» подробности кода блокировки в вашем API.
// Конечные пользователи вашей структуры не должны заботиться о том, как она делает
// внутреннюю синхронизацию.
// Определяя структуру, в которой мьютекс должен защищать одно или больше значений,
// помещайте мьютекс выше тех полей, доступ к которым, он будет защищать.
type InMemory struct {
	lock sync.Mutex        // <-- этот мьютекс защищает
	m    map[string][]byte // <-- это поле под ним
	cfg  *config.Config
}

// NewInMemory слой взаимодействия с хранилищем в памяти
func NewInMemory(cfg *config.Config) *InMemory {
	return &InMemory{
		cfg: cfg,
		m:   make(map[string][]byte),
	}
}

func (s *InMemory) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	s.lock.Lock()
	sh2 := entity.Shorturl{}
	defer s.lock.Unlock()
	if err := json.Unmarshal(s.m[sh.Slug], &sh2); err != nil {
		panic(err)
	}
	if sh2.URL != "" {
		return &sh2, nil
	}
	return nil, ErrNotFound
}

func (s *InMemory) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	//s.lock.Lock()
	//defer s.lock.Unlock()
	//sh := entity.Shorturl{}
	//lst := entity.List{}
	//err := json.Unmarshal(s.m[slug], &sh)
	//	if err != nil {
	//		i.r.file.Seek(0, 0)
	//	}
	//	if sh.UserID == u.UserID {
	//		lst.URL = sh.URL
	//		lst.Slug = scripts.GetHost(s.cfg.HTTP, sh.Slug)
	//		u.Urls = append(u.Urls, lst)
	//	}
	//fmt.Printf("DD:: %s, %v, %s, %s", u.Slug, u.Del, u.UserID, u.URL)
	return nil, ErrNotFound
}
func (s *InMemory) Delete(ctx context.Context, u *entity.User) error {
	s.lock.Lock()
	var sh2 entity.Shorturl
	defer s.lock.Unlock()
	fmt.Printf("Count: %d\n", len(s.m))
	if len(s.m) < 1 {
		return nil
	}
	for _, slug := range u.DelLink {
		if err := json.Unmarshal(s.m[slug], &sh2); err != nil {
			//fmt.Printf("delete error in memory unmarshal: %e", err)
			continue
		}
		sh2.Del = true
		obj, err := json.Marshal(sh2)
		if err != nil {
			return fmt.Errorf("delete error in memory marshal: %e", err)
		}
		s.m[slug] = obj
	}
	//fmt.Printf("DD:: %s, %v, %s, %s", sh2.Slug, sh2.Del, sh2.UserID, sh2.URL)
	return nil
}
func (s *InMemory) Put(ctx context.Context, sh *entity.Shorturl) error {
	s.Post(ctx, sh)
	return nil
}

func (s *InMemory) Post(ctx context.Context, sh *entity.Shorturl) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.m[sh.Slug]; ok {
		return ErrAlreadyExists
	}

	obj, err := json.Marshal(sh)
	if err != nil {
		return ErrNotFound
	}
	s.m[sh.Slug] = obj
	//s.m[sh.Slug] = sh.URL
	return nil
}
