// Package shorturlerror
// Как наш хендлер узнает, что надо отдать ошибку например 404 подробнее:
// https://youtu.be/dtLj-BCBi6I?t=736

package shorturlerror

import (
	"net/http"

	"github.com/pkg/errors"
)

type shortHandler func(w http.ResponseWriter, r *http.Request) error

// Middleware (Связующее программное обеспечение)
func Middleware(h shortHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *ShorturlError
		err := h(w, r)
		if err != nil {
			st := r.Header.Get("Content-Type")
			w.Header().Set("Content-Type", st)
			// проверяем наша ошибка пришла или какая-то левая
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*ShorturlError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}
			w.WriteHeader(http.StatusTeapot)
			// получаю все системные ошибки обернуто
			w.Write(systemError(err).Marshal())
		}
	}
}
