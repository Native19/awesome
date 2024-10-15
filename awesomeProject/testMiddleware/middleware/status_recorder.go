package middleware

import "net/http"

type LogWriter interface {
	SetStatus(int)
	SetBool(bool)
}

type statusRecorder struct {
	http.ResponseWriter
	Status int
	IsTrue bool
}

func (r *statusRecorder) Unwrap() http.ResponseWriter {
	return r.ResponseWriter
}

func (r *statusRecorder) Write(content []byte) (int, error) {
	// Если WriteHeader не был вызван.
	if r.Status == 0 {
		r.WriteHeader(http.StatusOK)
	}
	return r.ResponseWriter.Write(content)
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.Status = statusCode
}
func (r *statusRecorder) SetStatus(status int) {
	r.Status = status
}

func (r *statusRecorder) SetBool(status bool) {
	r.IsTrue = status
}
