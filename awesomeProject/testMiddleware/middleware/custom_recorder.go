package middleware

import "net/http"

type CustomLog interface {
	SetError(err error)
	SetMsg(msg string)
}

type customRecorder struct {
	http.ResponseWriter
	status int
	Msg    string
	Err    error
}

func (cr *customRecorder) Unwrap() http.ResponseWriter {
	return cr.ResponseWriter
}

func (cr *customRecorder) WriteHeader(statusCode int) {
	cr.ResponseWriter.WriteHeader(statusCode)
	cr.status = statusCode
}

func (cr *customRecorder) SetError(err error) {
	cr.Err = err
}
func (cr *customRecorder) SetMsg(msg string) {
	cr.Msg = msg
}
