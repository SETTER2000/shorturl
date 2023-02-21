// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Shorturl -.
type Shorturl struct {
	Slug   string `json:"slug" example:"1674872720465761244B_5"`
	URL    string `json:"url" example:"https://example.com/go/to/home.html"`
	UserID string `json:"user_id,omitempty"`
}
type List struct {
	Slug string `json:"short_url" example:"1674872720465761244B_5"`
	URL  string `json:"original_url" example:"https://example.com/go/to/home.html"`
}
type User struct {
	UserID string `json:"user_id" example:"1674872720465761244B_5"`
	Urls   []List
}

type ShorturlResponse struct {
	URL string `json:"result"`
}
