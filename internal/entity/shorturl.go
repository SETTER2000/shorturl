// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Shorturl -.
type Shorturl struct {
	Slug string `json:"slug" example:"^(\d{19})(\w{3})$"`
	URL  string `json:"link" example:"^http|https://example.com/go/to/home.html"`
}
