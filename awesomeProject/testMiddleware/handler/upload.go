package handler

import (
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

	var log middleware.LogWriter
	for {

		try, ok := w.(middleware.LogWriter)
		if ok {
			log = try
			break
		}
		unW, ok := w.(middleware.Unwrapper)
		if ok {
			w = unW.Unwrap()
		} else {
			break
		}
	}

	if log != nil {
		log.SetBool(true)
		log.SetStatus(http.StatusBadGateway)
		fmt.Println("upload успешно присвоил true и 502")
	}
}
