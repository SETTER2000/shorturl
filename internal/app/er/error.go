package er

import "errors"

// ErrNotFound ошибка в случаи отсутствия данных
// ErrAlreadyExists ошибка в случаи если данные уже существуют
// ErrBadRequest ошибка в случаи не корректного запроса
// ErrAccessDenied ошибка в случаи отсутствия права доступа
// ErrForbidden ошибка доступа к запрошенному ресурсу в случаи когда доступ запрещен
// ErrUserIDRequired -.
// ErrStatusConflict -.
var (
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrBadRequest     = errors.New("bad request")
	ErrAccessDenied   = errors.New(`access denied`)
	ErrForbidden      = errors.New(`forbidden`)
	ErrUserIDRequired = errors.New("user id required")
	ErrStatusConflict = errors.New("status conflict")
	ErrStatusGone     = errors.New("status gone")
)

type response struct {
	Error string `json:"error" example:"message"`
}