package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type MetaLogger struct {
	http.Handler
	name   string
	status int
	err    error
}

func NewMetaLogger(name string) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return &MetaLogger{
			Handler: handler,
			name:    name,
		}
	}
}

func (ml *MetaLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	fmt.Println("logger start")

	recorder := statusRecorder{
		ResponseWriter: w,
	}

	defer func() {
		fmt.Println("logger stop")
		fmt.Printf("work time : %v\n", time.Since(t1))
		fmt.Printf("status code : %v\n", recorder.Status)
		fmt.Printf("is true? : %v\n", recorder.IsTrue)
	}()

	ml.Handler.ServeHTTP(&recorder, r)
}
