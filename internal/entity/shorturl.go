// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Shorturl -.
type Shorturl struct {
	Slug string `json:"-" example:"1674872720465761244B_5"`
	URL  string `json:"url" example:"https://example.com/go/to/home.html"`
}
