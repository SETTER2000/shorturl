// Package v1 реализует пути маршрутизации. Каждая служба в своем файле.
package v1

import (
	"fmt"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// NewRouter -.
// Swagger spec:
// @title       Shorturl
// @description URL shortener server
// @version     0.1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypePlainText))

	r.Route("/", func(r chi.Router) {
		r.Post("/", LongURL) // POST /
		r.Get("/{key}", ShortURL)
	})
}

type Link struct {
	Slug string `json:"slug"`
	URL  string `json:"link"`
}

var db = make(map[string]string)

func LongURL(w http.ResponseWriter, r *http.Request) {
	data := &Link{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(string(body))
	data.URL = string(body)
	shorturl, err := dbNewLink(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// TODO не знаю как сюда прокинуть конфиг, чтоб убрать весь hardcode
	w.Write([]byte(fmt.Sprintf("http://localhost:8080/" + shorturl)))
}

func dbNewLink(link *Link) (string, error) {
	link.Slug = scripts.UniqueString()
	db[link.Slug] = link.URL
	return fmt.Sprintf(link.Slug), nil
}

func urlListFunc() []string {
	var list []string
	for _, c := range db {
		list = append(list, c)
	}
	return list
}

func ShortURL(w http.ResponseWriter, r *http.Request) {
	shorturl := chi.URLParam(r, "key")
	if shorturl == "" {
		http.Error(w, "key param is missed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Add("Location", urlFunc(shorturl))
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// urlFunc — вспомогательная функция для вывода определённого URL.
func urlFunc(id string) string {
	if c, ok := db[id]; ok {
		return c
	}
	return ""
}
