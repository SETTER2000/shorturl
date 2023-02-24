// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "github.com/SETTER2000/shorturl/config"

// Shorturl -.
type Shorturl struct {
	Slug   string `json:"slug" example:"1674872720465761244B_5"`             // Строковый идентификатор
	URL    string `json:"url" example:"https://example.com/go/to/home.html"` // URL для сокращения
	UserID string `json:"user_id,omitempty"`
	*config.Config
}
type List struct {
	Slug string `json:"short_url" example:"1674872720465761244B_5"`                 // Строковый идентификатор
	URL  string `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
}
type User struct {
	UserID string `json:"user_id" example:"1674872720465761244B_5"`
	Urls   []List
}

type ShorturlResponse struct {
	URL string `json:"result"` // URL для сокращения
}

type Batch struct {
	Slug string `json:"correlation_id" example:"1674872720465761244B_5"`            // Строковый идентификатор
	URL  string `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
	//URL  string `json:"short_url" example:"https://example.com/go/to/home.html"`    // Результирующий сокращённый URL
}

type ShortenResponse struct {
	Slug string `json:"correlation_id" example:"1674872720465761244B_5"`            // Строковый идентификатор
	URL  string `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
}
