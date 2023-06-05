package repo

import (
	"context"
	"sync"

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
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
	cfg  *config.Config
	m    map[string]entity.Shorturls // <-- это поле под ним
	lock sync.Mutex                  // <-- этот мьютекс защищает
}

// NewInMemory слой взаимодействия с хранилищем в памяти.
func NewInMemory(cfg *config.Config) *InMemory {
	return &InMemory{
		cfg: cfg,
		m:   make(map[string]entity.Shorturls),
	}
}

// Get получить конкретный URL по идентификатору этого URL и
// только если этот линк записал текущий пользователь.
func (s *InMemory) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	u, err := s.searchBySlug(sh)
	if err != nil {
		return nil, ErrNotFound
	}
	return u, nil
}

// search by slug
func (s *InMemory) searchBySlug(sh *entity.Shorturl) (*entity.Shorturl, error) {
	shorts := entity.Shorturls{}
	for _, uid := range s.m {
		for j := 0; j < len(uid); j++ {
			shorts = append(shorts, uid[j])
		}
	}
	for _, short := range shorts {
		if short.Slug == sh.Slug {
			sh.URL = short.URL
			sh.UserID = short.UserID
			sh.Del = short.Del
			break
		}
	}
	return sh, nil
}

// GetAllUrls получить все URL
func (s *InMemory) GetAllUrls() (entity.CountURLs, error) {
	return entity.CountURLs(len(s.m)), nil
}

// GetAll получить все URL пользователя по идентификатору.
func (s *InMemory) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	return nil, ErrNotFound
}

// Put обновить данные в память.
func (s *InMemory) Put(ctx context.Context, sh *entity.Shorturl) error {
	ln := len(s.m[sh.UserID])
	if ln < 1 {
		s.Post(ctx, sh)
		return nil
	}
	for j := 0; j < ln; j++ {
		if s.m[sh.UserID][j].Slug == sh.Slug {
			s.m[sh.UserID][j].URL = sh.URL
			s.m[sh.UserID][j].Del = sh.Del
			return nil
		}
	}
	return s.Post(ctx, sh)
}

// Post сохранить данные в память.
//
//	{
//		UserID: ShortURL{
//			UserID: str1
//		}
//	}
func (s *InMemory) Post(ctx context.Context, sh *entity.Shorturl) error {
	s.m[sh.UserID] = append(s.m[sh.UserID], *sh)
	return nil
}

// Delete - удаляет URLы переданный в запросе, только если есть права данного пользователя.
func (s *InMemory) Delete(ctx context.Context, u *entity.User) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.delete(u)
}

func (s *InMemory) delete(u *entity.User) error {
	for j := 0; j < len(s.m[u.UserID]); j++ {
		for _, slug := range u.DelLink {
			if s.m[u.UserID][j].Slug == slug {
				// изменяет флаг del на true, в результате url
				// становиться недоступным для пользователя
				s.m[u.UserID][j].Del = true
			}
		}
	}
	return nil
}

// Read - читает данные из памяти.
func (s *InMemory) Read() error {
	return nil
}

// Save - сохраняет данные в памяти.
func (s *InMemory) Save() error {
	return nil
}
