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

	// Shorturl хранит параметры URL.
	Shorturl struct {
		Slug           string `json:"slug,omitempty" example:"1674872720465761244B_5"`             // Строковый идентификатор
		URL            string `json:"url,omitempty" example:"https://example.com/go/to/home.html"` // URL для сокращения
		*config.Config `json:"-"`
		UserID         string `json:"user_id,omitempty"`
		Del            bool   `json:"del"`
	}

	// List -.
	List struct {
		Slug string `json:"short_url" example:"1674872720465761244B_5"`                 // Строковый идентификатор
		URL  string `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
	}

	// User -.
	User struct {
		UserID  string `json:"user_id" example:"1674872720465761244B_5"`
		DelLink []string
		Urls    []List
	}

	// ShorturlResponse -.
	ShorturlResponse struct {
		URL string `json:"result"` // URL для сокращения
	}

	// Batch -.
	Batch struct {
		Slug string `json:"correlation_id" example:"1674872720465761244B_5"`            // Строковый идентификатор
		URL  string `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
	}

	// ShortenResponse -.
	ShortenResponse struct {
		Slug string `json:"correlation_id" example:"1674872720465761244B_5"`        // Строковый идентификатор
		URL  string `json:"short_url" example:"https://example.com/correlation_id"` // URL для сокращения
	}

	// Short -.
	Short interface{}
)
