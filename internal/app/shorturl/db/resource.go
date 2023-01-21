package db

import (
	"context"
	"fmt"
	"github.com/SETTER2000/shorturl/internal/app/shorturl"
)

type repository struct {
	url map[string]string
}

func (d *repository) Create(ctx context.Context, shorturl shorturl.ShortUrl) (string, error) {
	if shorturl.Key == "" {
		return "", fmt.Errorf("short url creation error, key is empty")
	}
	d.url[shorturl.Key] = shorturl.Url
	//fmt.Printf("Размер хранилища: %v\n", len(d.url))
	//fmt.Printf("Ключ хранилища: %s\n", shorturl.Key)
	//fmt.Printf("Сейчас в хранилище: %s\n", d.url[shorturl.Key])
	return "http://localhost:8080?q=" + shorturl.Key, nil
}
func (d *repository) FindOne(ctx context.Context, id string) (s string, err error) {
	if id == "" {
		return "", fmt.Errorf("id request empty")
	}
	if d.url[id] == "" {
		return "", fmt.Errorf("not found")
	}
	longUrl := d.url[id]
	return longUrl, nil
}
func NewRepository() shorturl.Repository {
	r := &repository{}
	r.url = make(map[string]string)
	return r
}
