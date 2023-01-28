// Package v1 реализует пути маршрутизации. Каждая служба в своем файле.
package v1

import (
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// NewRouter -.
// Swagger spec:
// @title       Shorturl
// @description URL shortener server
// @version     0.1.0
// @host        localhost:8080
// @BasePath    /
func NewRouter(handler *chi.Mux, l logger.Interface, s usecase.Shorturl) {
	handler.Use(middleware.RequestID)
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)
	handler.Use(middleware.URLFormat)
	handler.Use(render.SetContentType(render.ContentTypePlainText))

	//handler.Get("/status", StatusHandler)

	// Routers
	h := handler.Route("/", func(r chi.Router) {
		r.Routes()
	})
	{
		newShorturlRoutes(h, s, l)
	}
}

//func StatusHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	// намеренно сделана ошибка в JSON
//	w.Write([]byte(`{"status":"ok"}`))
//}
//
//type Link struct {
//	Slug string `json:"slug"`
//	URL  string `json:"link"`
//}
//
//var db = make(map[string]string)
//
//// LongURL принимает длинный URL и возвращает короткий
//func LongURL(w http.ResponseWriter, r *http.Request) {
//	data := &Link{}
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	fmt.Println(string(body))
//	data.URL = string(body)
//	shorturl, err := dbNewLink(data)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	w.Header().Set("Content-Type", "text/plain")
//	w.WriteHeader(http.StatusCreated)
//	w.Write([]byte(fmt.Sprintf("http://localhost:8080/" + shorturl)))
//}
//
//func ShortURL(w http.ResponseWriter, r *http.Request) {
//	shorturl := chi.URLParam(r, "key")
//	if shorturl == "" {
//		http.Error(w, "key param is missed", http.StatusBadRequest)
//		return
//	}
//
//	w.Header().Set("Content-Type", "text/plain")
//	w.Header().Add("Location", urlFunc(shorturl))
//	w.WriteHeader(http.StatusTemporaryRedirect)
//}

//func dbNewLink(link *Link) (string, error) {
//	link.Slug = scripts.UniqueString()
//	db[link.Slug] = link.URL
//	return fmt.Sprintf(link.Slug), nil
//}

//func urlListFunc() []string {
//	var list []string
//	for _, c := range db {
//		list = append(list, c)
//	}
//	return list
//}

// urlFunc — вспомогательная функция для вывода определённого URL.
//func urlFunc(id string) string {
//	if c, ok := db[id]; ok {
//		return c
//	}
//	return ""
//}
