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

func (i *InMemory) Get(ctx context.Context, sh *entity.Shorturl) (*entity.Shorturl, error) {
	i.lock.Lock()
	sh2 := entity.Shorturl{}
	defer i.lock.Unlock()
	if err := json.Unmarshal(i.m[sh.Slug], &sh2); err != nil {
		panic(err)
	}
	if sh2.URL != "" {
		return &sh2, nil
	}
	return nil, ErrNotFound
}

func (i *InMemory) GetAll(ctx context.Context, u *entity.User) (*entity.User, error) {
	return nil, ErrNotFound
}
func (i *InMemory) Delete(ctx context.Context, u *entity.User) error {
	i.lock.Lock()
	var sh2 entity.Shorturl
	defer i.lock.Unlock()
	if len(i.m) < 1 {
		return nil
	}
	for _, slug := range u.DelLink {
		if err := json.Unmarshal(i.m[slug], &sh2); err != nil {
			continue
		}
		sh2.Del = true
		obj, err := json.Marshal(sh2)
		if err != nil {
			return fmt.Errorf("delete error in memory marshal: %e", err)
		}
		i.m[slug] = obj
	}
	return nil
}
func (i *InMemory) Put(ctx context.Context, sh *entity.Shorturl) error {
	return i.Post(ctx, sh)
}

func (i *InMemory) Post(ctx context.Context, sh *entity.Shorturl) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	if _, ok := i.m[sh.Slug]; ok {
		return ErrAlreadyExists
	}

	obj, err := json.Marshal(sh)
	if err != nil {
		return ErrNotFound
	}
	i.m[sh.Slug] = obj
	return nil
}
func (i *InMemory) Read() error {
	return nil
}
func (i *InMemory) Save() error {
	return nil
}
