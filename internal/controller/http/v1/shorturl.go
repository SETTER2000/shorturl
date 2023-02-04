package v1

import (
	"encoding/json"
	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/SETTER2000/shorturl/scripts"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type shorturlRoutes struct {
	s   usecase.Shorturl
	l   logger.Interface
	cfg config.HTTP
}

func newShorturlRoutes(handler chi.Router, s usecase.Shorturl, l logger.Interface, cfg config.HTTP) {
	sr := &shorturlRoutes{s, l, cfg}

	handler.Group(func(r chi.Router) {
		r.Post("/{some_url}", sr.shorten) // POST /
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
	shorturl, err := r.s.ShortLink(res, req)
	if err != nil {
		r.l.Error(err, "http - v1 - shortLink")
		http.Error(res, "key param is missed", http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.Header().Add("Location", shorturl)
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
	data := entity.Shorturl{}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	data.URL = string(body)
	shorturl, err := r.s.LongLink(data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(scripts.GetHost(r.cfg, shorturl)))
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
	shorturl, err := r.s.Shorten(data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	obj, err := json.Marshal(shorturlResponse{scripts.GetHost(r.cfg, shorturl)})
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(obj)
}
