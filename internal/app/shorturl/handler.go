package shorturl

import (
	"context"
	"fmt"
	"github.com/SETTER2000/shorturl/internal/app/handlers"
	"github.com/SETTER2000/shorturl/internal/app/shorturlerror"
	"github.com/go-chi/chi/v5"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const (
	searchURL = "/"
)

// Handler - все обработчики запросов в одной структуре реализующей интерфейс,
// интерфейс с единой точкой входа в виде роутера
type handler struct {
	repository Repository
}

func NewHandler(repository Repository) handlers.Handler {
	return &handler{
		repository: repository,
	}
}

// Register обязательно реализуем этот интерфейсный метод в структуре handler
func (h *handler) Register(router *chi.Mux) {
	// добавить длинный url и получить в ответ короткий
	router.Post(searchURL, shorturlerror.Middleware(h.AddUrl))
	// получить ресурс по короткому url
	router.Get(searchURL, shorturlerror.Middleware(h.GetShortUrl))
}

func (h *handler) GetShortUrl(w http.ResponseWriter, r *http.Request) error {
	key := r.URL.Query().Get("query")
	//key := chi.URLParam(r, "q")
	//key = strings.Split(key, "=")[1]
	//key := r.URL.RequestURI()
	shorturl, err := h.repository.FindOne(context.TODO(), key)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return shorturlerror.ErrNotFound
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Add("Location", shorturl)
	w.WriteHeader(http.StatusTemporaryRedirect)
	//w.Write([]byte(shorturl))
	return nil
}

// AddUrl добавить url в db (ну типа db)
func (h *handler) AddUrl(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	var url ShortUrl
	url.URL = string(body)
	url.Key = genStr()
	shorturl, err := h.repository.Create(context.TODO(), url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("http://localhost:8080?query=" + shorturl)))
	return nil
}
func genStr() string {
	// generate string
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	//specials := "~=+%^*/()[]{}/!@#$?|"
	specials := "_"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + digits + specials
	length := 8
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}
