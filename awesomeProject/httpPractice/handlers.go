package httpPractice

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Handler1 struct{}

func (h Handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	_, err := w.Write([]byte("Hello handler1"))
	if err != nil {
		fmt.Println(err)
	}
}

type Handler2 struct{}

func (h Handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	_, err := w.Write([]byte("Hello handler2"))
	if err != nil {
		fmt.Println(err)
	}
	errCh, ok := r.Context().Value("errCh").(chan<- error)
	if !ok {
		fmt.Println(ok)
		return
	}
	errCh <- errors.New("ошибка handler2")
	time.Sleep(20 * time.Second)
}

type Handler3 struct{}

func (h Handler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()
	str := r.URL
	_, err := w.Write([]byte(fmt.Sprintf("Hello handler3 id: %v", str)))
	if err != nil {
		fmt.Println(err)
	}
}
