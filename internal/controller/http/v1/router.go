// Package v1 реализует пути маршрутизации. Каждая служба в своем файле.
package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/SETTER2000/shorturl/config"
	"github.com/SETTER2000/shorturl/internal/usecase"
	"github.com/SETTER2000/shorturl/internal/usecase/encryp"
	"github.com/SETTER2000/shorturl/pkg/compress/gzip"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Shorturl
// @description URL shortener server
// @version     0.1.0
// @host        localhost:8080
// @BasePath    /
func NewRouter(handler *chi.Mux, l logger.Interface, s usecase.Shorturl, cfg *config.Config) {
	headerTypes := []string{
		"application/javascript",
		"application/x-gzip",
		"application/gzip",
		"application/json",
		"text/css",
		"text/html",
		"text/plain",
		"text/xml",
	}
	// AllowContentType применяет белый список запросов Content-Types,
	// в противном случае отвечает статусом 415 Unsupported Media Type.
	handler.Use(middleware.AllowContentType(headerTypes...))
	handler.Use(middleware.Compress(5, headerTypes...))
	handler.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	handler.Use(middleware.RequestID)
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)
	handler.Use(render.SetContentType(render.ContentTypePlainText))
	handler.Use(encryp.EncryptionCookie(cfg))
	handler.Use(gzip.DeCompressGzip)
	handler.Mount("/debug", middleware.Profiler())

	// Swagger
	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	sr := &shorturlRoutes{s, l, cfg}

	handler.Route("/", func(handler chi.Router) {
		//handler.Handle("/", gzip.DeCompressGzip(http.HandlerFunc(sr.longLink)))
		handler.Post("/", sr.longLink)
		handler.Get("/{key}", sr.shortLink)
		handler.Get("/ping", sr.connect)
	})

	// Routers
	h := handler.Route("/api", func(r chi.Router) {
		r.Routes()
	})
	{
		newShorturlRoutes(h, s, l, cfg)
	}
}
