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
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	sh, err := r.s.ShortLink(res, req)
	if err != nil {
		r.l.Error(err, "http - v1 - shortLink")
		http.Error(res, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}
	//fmt.Printf("Content--TYPE3 string(body): ind: %s\n", res.Header().Get("Content-Type"))
	//fmt.Printf("Content--TYPE3 string(body): ind: %s\n", res.Header().Get("Content-Type"))
	fmt.Printf("Content--TYPE4 string(body):  %s\n", http.DetectContentType(body))
	res.Header().Set("Content-Type", http.DetectContentType(body))
	//res.Header().Set("Content-Encoding", "gzip")
	//res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//res.Header().Set("Content-Type", "application/gzip; charset=utf-8")
	res.Header().Add("Location", sh.URL)

	//log.Printf("HHH GET:::%v", res.Header())
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
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Content--TYPE string(body): ind: %s\n", res.Header().Get("Content-Type"))
	fmt.Printf("Content--TYPE2 string(body): ind: %s\n", http.DetectContentType(body))
	fmt.Printf("POST string(body): ind: %v", body)
	fmt.Printf("POST string(body): ind: %s", body)
	data := entity.Shorturl{}
	data.URL = string(body)
	//fmt.Println(matched, err)
	shorturl, err := r.s.LongLink(&data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	d := scripts.GetHost(r.cfg, shorturl)
	res.Header().Set("Content-Type", http.DetectContentType(body))
	//res.Header().Set("Content-Type", "text/plain; charset=utf-8")
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
	shorturl, err := r.s.Shorten(&data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	respURL := scripts.GetHost(r.cfg, shorturl)
	obj, err := json.Marshal(shorturlResponse{respURL})
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	res.Write(obj)
}
