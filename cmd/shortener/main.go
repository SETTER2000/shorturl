package main

import (
	"io"
	"log"
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
	Url string
}

var Link = Urls{}

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
		Link.Url = string(b)
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
		w.Header().Add("Location", Link.Url)
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Write([]byte(Link.Url))

	default:
		http.Error(w, "Only POST|GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	http.HandleFunc(rootDir, Handlers)
	log.Fatal(http.ListenAndServe(SERVER, nil))
}
