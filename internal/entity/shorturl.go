// Package entity определяет основные сущности для бизнес-логики (сервиса), сопоставления базы данных и объектов ответа HTTP, если они подходят.
package entity

import "github.com/SETTER2000/shorturl/config"

// Type - содержит все сущности используемые сервисом.
type (

	// CorrelationOrigin -.
	CorrelationOrigin []Batch

	// Response -.
	Response []ShortenResponse

	// Shorturls -.
	Shorturls []Shorturl

	// Slug -.
	Slug string

	// UserID -.
	UserID string

	// URL -.
	URL string

	// Shorturl хранит параметры URL.
	Shorturl struct {
		Slug           `json:"slug,omitempty" example:"1674872720465761244B_5"`             // Строковый идентификатор
		URL            `json:"url,omitempty" example:"https://example.com/go/to/home.html"` // URL для сокращения
		*config.Config `json:"-"`
		UserID         `json:"user_id,omitempty"`
		Del            bool `json:"del"`
	}

	// List -.
	List struct {
		ShortURL URL                                                                 `json:"short_url" example:"1674872720465761244B_5"` // Строковый идентификатор
		URL      `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
	}

	// User -.
	User struct {
		UserID  `json:"user_id" example:"1674872720465761244B_5"`
		DelLink []Slug
		Urls    []List
	}

	// ShorturlResponse -.
	ShorturlResponse struct {
		URL `json:"result"` // URL для сокращения
	}

	// Batch -.
	Batch struct {
		Slug `json:"correlation_id" example:"1674872720465761244B_5"`            // Строковый идентификатор
		URL  `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
	}

	// ShortenResponse -.
	ShortenResponse struct {
		Slug `json:"correlation_id" example:"1674872720465761244B_5"`        // Строковый идентификатор
		URL  `json:"short_url" example:"https://example.com/correlation_id"` // URL для сокращения
	}

	// Short -.
	Short interface{}

	// CountUsers кол-во пользователей в сервисе
	CountUsers int

	// CountURLs кол-во сокращённых URL в сервисе
	CountURLs int

	// Static -.
	Static struct {
		CountURLs  `json:"urls,omitempty"`  // кол-во сокращённых URL в сервисе
		CountUsers `json:"users,omitempty"` // кол-во пользователей в сервисе
	}
)
