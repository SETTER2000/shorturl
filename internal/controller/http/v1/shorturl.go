package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

	handler.Group(func(r chi.Router) {
		r.Post("/{some_url}", sr.shorten) // POST /
	})

	handler.Route("/user", func(r chi.Router) {
		r.Get("/urls", sr.urls)
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
	sh, err := r.s.ShortLink(res, req)
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
func (r *shorturlRoutes) ping(res http.ResponseWriter, req *http.Request) {
	dsn, ok := os.LookupEnv("DATABASE_DSN")
	if !ok || dsn == "" {
		dsn = r.cfg.Storage.ConnectDB
		if dsn == "" {
			fmt.Printf("connect DSN string is empty: %v\n", dsn)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_DSN"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			//os.Exit(1)
			res.WriteHeader(http.StatusInternalServerError)
		}
		defer conn.Close(context.Background())
		//var name string
		//var weight int64
		//err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		//	os.Exit(1)
		//}

		//fmt.Println(name, weight)
		//fmt.Printf("connect... %s\n", dsn)
		fmt.Printf("connect... \n")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("connect... "))
		//res.Write([]byte(fmt.Sprintf("connect... %s\n", dsn)))
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
	// при чтении вернётся распакованный слайс байт
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	data := entity.Shorturl{}
	data.URL = string(body)
	data.UserID = req.Context().Value("access_token").(string)
	shorturl, err := r.s.LongLink(&data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	d := scripts.GetHost(r.cfg.HTTP, shorturl)
	res.Header().Set("Content-Type", http.DetectContentType(body))
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(d))
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
	data := entity.Shorturl{}
	resp := entity.ShorturlResponse{}
	body, err := io.ReadAll(req.Body)
	fmt.Printf("DUDA 55 http shorten:: %v\n", string(body))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	fmt.Printf("DUDA http shorten:: %v\n", data)
	data.UserID = req.Context().Value(r.cfg.Cookie.AccessTokenName).(string)
	shorturl, err := r.s.Shorten(&data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	resp.URL = scripts.GetHost(r.cfg.HTTP, shorturl)
	obj, err := json.Marshal(resp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(obj)
}

// GET
func (r *shorturlRoutes) urls(res http.ResponseWriter, req *http.Request) {
	u := entity.User{}
	userID := req.Context().Value("access_token")
	if userID == nil {
		res.Write([]byte(fmt.Sprintf("Not access_token and user_id: %s", userID)))
	}
	u.UserID = fmt.Sprintf("%s", userID)
	user, err := r.s.UserAllLink(&u)
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
