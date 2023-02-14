package v1

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
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
	"strings"
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
	//b, err := Decompress(sh.URL)
	//if err != nil {
	//	fmt.Errorf("errrr %e", err)
	//}
	//fmt.Printf("Content--TYPE3 string(body): ind: %s\n", res.Header().Get("Content-Type"))
	fmt.Printf("DetectContentType в GET: %s\n", http.DetectContentType(body))
	//fmt.Printf("Content--sh.URL:  %s\n", sh.URL)
	res.Header().Set("Content-Type", http.DetectContentType(body))
	res.Header().Add("Content-Encoding", "gzip")
	//res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//res.Header().Set("Content-Type", "application/gzip; charset=utf-8")
	//res.Header().Add("Accept-Encoding", "application/json")
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

	if !strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		fmt.Println("если gzip не поддерживается, передаём управление дальше без изменений")
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Accept-Encoding от клиента:%v\n", res.Header().Get("Accept-Encoding"))
	fmt.Printf("Content-Encoding от клиента:%v\n", res.Header().Get("Content-Encoding"))
	//b, err := Decompress(body)
	//if err != nil {
	//	fmt.Printf("errrr %e\n", err)
	//}
	//LengthHandle(res, req)
	//fmt.Printf("Accept-Encoding от клиента:%v\n", res.Header().Get("Accept-Encoding"))
	//fmt.Printf("Content--TYPE string(body): ind: %s\n", res.Header().Get("Content-Type"))
	//fmt.Printf("Content--TYPE2 string(body): ind: %s\n", http.DetectContentType(body))
	//fmt.Printf("POST string(bytes):%v\n", body)
	//fmt.Printf("POST string(body):%s\n", body)

	data := entity.Shorturl{}
	data.URL = string(body)
	fmt.Printf("SPLIT string(body):%s\n", strings.Split(data.URL, "\n"))
	fmt.Printf("DATA  в POST::%s\n", data)
	fmt.Printf("DATA URL в POST::%s\n", data.URL)
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

// LengthHandle возвращает размер распакованных данных.
func (r *shorturlRoutes) longLink3(res http.ResponseWriter, req *http.Request) {
	// создаём *gzip.Reader, который будет читать тело запроса
	// и распаковывать его
	gz, err := gzip.NewReader(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// не забывайте потом закрыть *gzip.Reader
	defer gz.Close()

	// при чтении вернётся распакованный слайс байт
	body, err := io.ReadAll(gz)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Fprintf(res, "Length: %d", len(body))

	data := entity.Shorturl{}
	data.URL = string(body)
	fmt.Printf("SPLIT string(body):%s\n", strings.Split(data.URL, "\n"))
	fmt.Printf("DATA  в POST::%s\n", data)
	fmt.Printf("DATA URL в POST::%s\n", data.URL)
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
func Decompress(data []byte) ([]byte, error) {
	// переменная r будет читать входящие данные и распаковывать их
	r := flate.NewReader(bytes.NewReader(data))
	defer r.Close()

	var b bytes.Buffer
	// в переменную b записываются распакованные данные
	_, err := b.ReadFrom(r)
	if err != nil {
		return nil, fmt.Errorf("failed decompress data: %v", err)
	}

	return b.Bytes(), nil
}
func (r *shorturlRoutes) longLink2(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	body, err = Decompress(body)
	if err != nil {
		fmt.Printf("errrr %e", err)
	}

	fmt.Printf("Content==string(Content-Type):%s\n", res.Header().Get("Content-Type"))
	fmt.Printf("Content== string(body):%s\n", http.DetectContentType(body))
	fmt.Printf("POST== string(bytes):%v\n", body)
	fmt.Printf("POST== string(body):%s\n", body)
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
