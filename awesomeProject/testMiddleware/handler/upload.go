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

	log, err := middleware.GetWriter[middleware.LogWriter](w)
	if err == nil {
		log.SetBool(true)
		log.SetStatus(http.StatusBadGateway)
		fmt.Println("upload успешно присвоил true и 502")
	}

	observer, err := middleware.GetWriter[middleware.CustomLog](w)
	if err == nil {
		observer.SetError(errors.New("какая-то ошибка"))
		observer.SetMsg("данные обсервера")
		fmt.Println("upload успешно присвоил ошибку и msg")
	}
}
