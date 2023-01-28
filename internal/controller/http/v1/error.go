package v1

import (
	"github.com/go-chi/chi/v5"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(c *chi.Context, code int, msg string) {
	return
}
