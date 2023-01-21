package shorturlerror

import (
	"encoding/json"
)

var ErrNotFound = NewShorturlError(
	nil, "not found", "", "NF-000004")

// ShorturlError собственная система ошибок,
// реализован интерфейсный метод Error, чтоб соответствовать
type ShorturlError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func (e *ShorturlError) Error() string {
	return e.Message
}

// Unwrap - чтобы соответствовать интерфейсному методу
func (e *ShorturlError) Unwrap() error { return e.Err }

func (e *ShorturlError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func NewShorturlError(
	err error, message, developerMessage, code string) *ShorturlError {
	return &ShorturlError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func systemError(err error) *ShorturlError {
	return NewShorturlError(
		err, "internal system error", err.Error(), "SH-000000")
}
