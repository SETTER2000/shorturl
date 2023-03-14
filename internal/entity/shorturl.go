// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "github.com/SETTER2000/shorturl/config"

// CorrelationOrigin -.
type CorrelationOrigin []Batch
type Response []ShortenResponse
type Shorturls []Shorturl

// Shorturl -.
type Shorturl struct {
	Slug           string `json:"slug,omitempty" example:"1674872720465761244B_5"`             // Строковый идентификатор
	URL            string `json:"url,omitempty" example:"https://example.com/go/to/home.html"` // URL для сокращения
	UserID         string `json:"user_id,omitempty"`
	Del            bool   `json:"del"`
	*config.Config `json:"-"`
	//*CorrelationOrigin `json:"correlation_origin,omitempty"`
}
type List struct {
	Slug string `json:"short_url" example:"1674872720465761244B_5"`                 // Строковый идентификатор
	URL  string `json:"original_url" example:"https://example.com/go/to/home.html"` // URL для сокращения
}
type User struct {
	UserID  string `json:"user_id" example:"1674872720465761244B_5"`
	Urls    []List
	DelLink []string
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
	Slug string `json:"correlation_id" example:"1674872720465761244B_5"`        // Строковый идентификатор
	URL  string `json:"short_url" example:"https://example.com/correlation_id"` // URL для сокращения
}

type Short interface{}
