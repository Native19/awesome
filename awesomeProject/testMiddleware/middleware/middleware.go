package middleware

import (
	"errors"
	"net/http"
)

func GetWriter[T any](w http.ResponseWriter) (writer T, err error) {
	for {
		try, ok := w.(T)
		if ok {
			return try, nil
		}
		unW, ok := w.(interface {
			Unwrap() http.ResponseWriter
		})
		if ok {
			w = unW.Unwrap()
		} else {
			break
		}
	}
	var t T
	return t, errors.New("не удалось извлечь тип")
}
