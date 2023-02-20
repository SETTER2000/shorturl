package v1

import (
	"encoding/json"
	"fmt"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

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
	data.UserId = req.Context().Value("access_token").(string)
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

type shorturlResponse struct {
	URL string `json:"result"`
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
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
	data.UserId = req.Context().Value(r.cfg.Cookie.AccessTokenName).(string)
	shorturl, err := r.s.Shorten(&data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	respURL := scripts.GetHost(r.cfg.HTTP, shorturl)
	obj, err := json.Marshal(shorturlResponse{respURL})
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
	userId := req.Context().Value("access_token")
	if userId == nil {
		res.Write([]byte(fmt.Sprintf("Not access_token and user_id: %s", userId)))
	}
	u.UserId = fmt.Sprintf("%s", userId)
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
