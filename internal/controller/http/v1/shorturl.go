package v1

import (
	"fmt"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"

	"github.com/SETTER2000/shorturl/internal/entity"
	"github.com/SETTER2000/shorturl/internal/usecase"
)

type shorturlRoutes struct {
	s usecase.Shorturl
	l logger.Interface
}

func newShorturlRoutes(handler chi.Router, s usecase.Shorturl, l logger.Interface) {
	sr := &shorturlRoutes{s, l}

	handler.Group(func(r chi.Router) {
		r.Get("/{key}", sr.shortLink) // GET /{key}
		r.Post("/", sr.longLink)      // POST /
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
// @ID          shortLink
// @Tags  	    shorturl
// @Accept      text
// @Produce     text
// @Success     307 {object} string
// @Failure     500 {object} response
// @Router      /{key} [get]
func (r *shorturlRoutes) longLink(res http.ResponseWriter, req *http.Request) {
	data := entity.Shorturl{}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(string(body))
	data.URL = string(body)
	shorturl, err := r.s.LongLink(data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(fmt.Sprintf("http://localhost:8080/" + shorturl)))
}
