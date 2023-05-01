package v1

import "errors"

// ErrNotFound ошибка в случаи отсутствия данных
// ErrAlreadyExists ошибка в случаи если данные уже существуют
// ErrBadRequest ошибка в случаи не корректного запроса
// ErrAccessDenied ошибка в случаи отсутствия права доступа
var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
	ErrBadRequest    = errors.New("bad request")
	ErrAccessDenied  = errors.New(`access denied`)
)

type response struct {
	Error string `json:"error" example:"message"`
}
