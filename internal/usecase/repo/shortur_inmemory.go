package repo

import "sync"

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

func NewInMemory() *InMemory {
	return &InMemory{
		m: make(map[string]string),
	}
}

func (s *InMemory) Get(key string) (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if v, ok := s.m[key]; ok {
		return v, nil
	}
	return "", ErrNotFound
}

func (s *InMemory) Put(key string, value string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.m[key]; ok {
		return ErrAlreadyExists
	}
	s.m[key] = value
	return nil
}