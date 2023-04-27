package encryp

import "errors"

// ErrNotFound ошибка в случаи отсутствия данных
// ErrAlreadyExists ошибка в случаи если данные уже существуют
// ErrBadRequest ошибка в случаи не корректного запроса
// ErrAccessDenied ошибка в случаи отстуствия права доступа
// ErrEncryptToken ошибка в случаи отстуствия ключа шифрования
var (
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("already exists")
	ErrBadRequest          = errors.New("bad request")
	ErrAccessDenied        = errors.New(`access denied`)
	ErrEncryptToken        = errors.New(`bad encryption keys`)
	ErrNewCipherNotCreated = errors.New(`NewCipher not created`)
	ErrNewGCMNotCreated    = errors.New(`NewGCM not created`)
)

type response struct {
	Error string `json:"error" example:"message"`
}
