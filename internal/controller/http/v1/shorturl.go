package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SETTER2000/shorturl/internal/usecase/repo"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type contextKey string

const userIDKey contextKey = "access_token"

type shorturlRoutes struct {
	s   usecase.Shorturl
	l   logger.Interface
	cfg *config.Config
}

func newShorturlRoutes(handler chi.Router, s usecase.Shorturl, l logger.Interface, cfg *config.Config) {
	sr := &shorturlRoutes{s, l, cfg}
	handler.Route("/user", func(r chi.Router) {
		r.Get("/urls", sr.urls)
	})
	handler.Route("/shorten", func(r chi.Router) {
		r.Post("/", sr.shorten) // POST /
		r.Post("/batch", sr.batch)
	})
}

// @Summary     Return short URL
// @Description Redirect to long URL
// @ID          shortLink
// @Tags  	    shorturl
// @Accept      text
// @Produce     text
// @Success     307 {object} string
// @Failure     500 {object} response
// @Router      /{key} [get]
func (r *shorturlRoutes) shortLink(res http.ResponseWriter, req *http.Request) {
	shorturl := chi.URLParam(req, "key")
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	data := entity.Shorturl{Config: r.cfg}
	data.Slug = shorturl
	sh, err := r.s.ShortLink(ctx, &data)
	if err != nil {
		r.l.Error(err, "http - v1 - shortLink")
		http.Error(res, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "text/plain")
	res.Header().Add("Content-Encoding", "gzip")
	res.Header().Add("Location", sh.URL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

// GET /ping, который при запросе проверяет соединение с базой данных
// при успешной проверке хендлер должен вернуть HTTP-статус 200 OK
// при неуспешной — 500 Internal Server Error
func (r *shorturlRoutes) connect(res http.ResponseWriter, req *http.Request) {
	dsn, ok := os.LookupEnv("DATABASE_DSN")
	if !ok || dsn == "" {
		dsn = r.cfg.Storage.ConnectDB
		if dsn == "" {
			fmt.Printf("connect DSN string is empty: %v\n", dsn)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_DSN"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		defer db.Close(context.Background())

		fmt.Printf("connect... \n")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("connect... "))
	}
}

// @Summary     Return short URL
// @Description Redirect to long URL
// @ID          longLink
// @Tags  	    shorturl
// @Accept      text
// @Produce     text
// @Success     201 {object} string
// @Failure     500 {object} response
// @Router      / [post]
func (r *shorturlRoutes) longLink(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	data := entity.Shorturl{Config: r.cfg}
	data.URL = string(body)
	data.Slug = scripts.UniqueString()
	//data.UserID = req.Context().Value("access_token").(string)
	shorturl, err := r.s.LongLink(ctx, &data)
	if err != nil {
		if errors.Is(err, repo.ErrAlreadyExists) {
			data2 := entity.Shorturl{Config: r.cfg}
			data2.URL = data.URL
			sh, err := r.s.ShortLink(ctx, &data2)
			if err != nil {
				r.l.Error(err, "http - v2 - shortLink")
				http.Error(res, fmt.Sprintf("%v", err), http.StatusBadRequest)
				return
			}
			shorturl = sh.Slug
			res.Header().Set("Content-Type", http.DetectContentType(body))
			res.WriteHeader(http.StatusConflict)
		} else {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
	d := scripts.GetHost(r.cfg.HTTP, shorturl)
	res.Header().Set("Content-Type", http.DetectContentType(body))
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(d))
}

// GET
func (r *shorturlRoutes) urls(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	u := entity.User{}
	userID := req.Context().Value("access_token")
	if userID == nil {
		res.Write([]byte(fmt.Sprintf("Not access_token and user_id: %s", userID)))
	}
	u.UserID = fmt.Sprintf("%s", userID)
	user, err := r.s.UserAllLink(ctx, &u)
	if err != nil {
		r.l.Error(err, "http - v1 - shortLink")
		http.Error(res, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	obj, err := json.Marshal(user.Urls)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%v", len(obj))
	res.Header().Set("Content-Type", "application/json")
	if string(obj) == "null" {
		res.WriteHeader(http.StatusNoContent)
	} else {
		res.WriteHeader(http.StatusOK)
	}
	res.Write(obj)
}

// @Summary     Return JSON short URL
// @Description Redirect to long URL
// @ID          shorten
// @Tags  	    shorturl
// @Accept      json
// @Produce     json
// @Success     307 {object} string
// @Failure     500 {object} response
// @Router      /{shorten} [post]
func (r *shorturlRoutes) shorten(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	data := entity.Shorturl{Config: r.cfg}
	resp := entity.ShorturlResponse{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	data.Slug = scripts.UniqueString()
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	//data.UserID = req.Context().Value(r.cfg.Cookie.AccessTokenName).(string)
	resp.URL, err = r.s.Shorten(ctx, &data)
	if err != nil {
		if errors.Is(err, repo.ErrAlreadyExists) {
			data2 := entity.Shorturl{Config: r.cfg}
			data2.URL = data.URL
			sh, err := r.s.ShortLink(ctx, &data2)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
			resp.URL = sh.Slug
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusConflict)
		} else {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
	resp.URL = scripts.GetHost(r.cfg.HTTP, resp.URL)
	obj, err := json.Marshal(resp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(obj)
}

func (r *shorturlRoutes) batch(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()
	data := entity.Shorturl{Config: r.cfg}
	CorrelationOrigin := entity.CorrelationOrigin{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.Unmarshal(body, &CorrelationOrigin); err != nil {
		panic(err)
	}

	var rs entity.Response
	var sr entity.ShortenResponse
	for _, bt := range CorrelationOrigin {
		data.URL = bt.URL
		data.Slug = bt.Slug
		_, err = r.s.Shorten(ctx, &data)
		if err != nil {
			if errors.Is(err, repo.ErrAlreadyExists) {
				res.WriteHeader(http.StatusConflict)
				return
			}
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		sr.Slug = data.Slug
		sr.URL = scripts.GetHost(r.cfg.HTTP, data.Slug)
		rs = append(rs, sr)
	}

	obj, err := json.Marshal(rs)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(obj)
}
