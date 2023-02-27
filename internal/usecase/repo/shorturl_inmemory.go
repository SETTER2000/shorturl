package repo

import (
	"context"
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
	m    map[string]string // <-- это поле под ним
}

// NewInMemory слой взаимодействия с хранилищем в памяти
func NewInMemory() *InMemory {
	return &InMemory{
		m: make(map[string]string),
	}
}

func (s *InMemory) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//sh := entity.Shorturl{}
	if v, ok := s.m[sh.Slug]; ok {
		//sh.Slug = key
		sh.URL = v
		return sh, nil
	}
	return nil, ErrNotFound
}

func (s *InMemory) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//sh := entity.Shorturl{}
	//if v, ok := s.m[key]; ok {
	//	sh.Slug = key
	//	sh.URL = v
	//	return &sh, nil
	//}
	return nil, ErrNotFound
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
	s.m[sh.Slug] = sh.URL
	return nil
}
