package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"awesome/awesomeProject/testMiddleware/middleware"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (u *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload")

	var log middleware.LogWriter

	if log != nil {
		log.SetBool(true)
		log.SetStatus(http.StatusBadGateway)
		fmt.Println("upload успешно присвоил true и 502")
	}
}

func getLog(log interface{}, w http.ResponseWriter) error {
	for {

		fmt.Println(reflect.TypeOf(log))
		typeL := reflect.TypeOf(log)

		try, ok := w.(typeL)
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
}
