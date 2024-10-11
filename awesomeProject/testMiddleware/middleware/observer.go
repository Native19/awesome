package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Observer struct {
	http.Handler
	name string
}

func NewObserver(name string) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return &Observer{
			Handler: handler,
			name:    name,
		}
	}
}

func (o *Observer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	recorder := customRecorder{
		ResponseWriter: w,
	}

	fmt.Println("observer start")
	defer func() {
		fmt.Println("observer stop")
		fmt.Println(recorder.Msg)
		fmt.Println(recorder.Err)
	}()

	o.Handler.ServeHTTP(&recorder, r)
}
