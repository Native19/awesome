package handler

import (
	"errors"
	"fmt"
	"net/http"

	"awesome/awesomeProject/testMiddleware/middleware"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (u *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload")

	log, err := getWriter[middleware.LogWriter](w)
	if err != nil {
		return
	}

	if log != nil {
		log.SetBool(true)
		log.SetStatus(http.StatusBadGateway)
		fmt.Println("upload успешно присвоил true и 502")
	}
}

func getWriter[T any](w http.ResponseWriter) (writer T, err error) {
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
