package handlers

import (
	"io"
	"net/http"
)

const (
	PORT    = ":8080"
	rootDir = "/"
	SERVER  = "localhost" + PORT
	DOMAIN  = "http://" + SERVER
)

type Urls struct {
	Key string
	URL string
}

var Link = Urls{}

func StatusHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	// намеренно сделана ошибка в JSON
	rw.Write([]byte(`{"status":"ok"}`))
}

func Handlers(w http.ResponseWriter, r *http.Request) {
	Link.Key = "ussr"
	urlResp := DOMAIN + "/search?query=" + Link.Key
	switch r.Method {
	case "POST":
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		Link.URL = string(b)
		ct := r.Header.Get("Content-Type")
		w.Header().Set("Content-Type", ct)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(urlResp))

	case "GET":
		QueryKey := r.URL.Query().Get("query")
		if QueryKey == "" {
			http.Error(w, "The query parameter is missing", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Add("Location", Link.URL)
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Write([]byte(Link.URL))

	default:
		http.Error(w, "Only POST|GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
}
