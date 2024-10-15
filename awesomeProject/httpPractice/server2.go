package httpPractice

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func Serv2() (*http.Server, chan error) {
	var errChan = make(chan error, 10)
	handlersErrCh := chan<- error(errChan)
	f := func(l net.Listener) context.Context {
		return context.WithValue(context.Background(), "errCh", handlersErrCh)
	}

	baseMux := http.NewServeMux()
	//baseMux.Handle("/", Handler2{})
	baseMux.Handle("/hand2", Handler2{})
	baseMux.HandleFunc("/hand1", Handler1{}.ServeHTTP)
	baseMux.HandleFunc("/id:5", Handler3{}.ServeHTTP)
	baseMux.HandleFunc("/custom", func(w http.ResponseWriter, r *http.Request) {
		defer func() { _ = r.Body.Close() }()
		_, err := w.Write([]byte("custom handler"))
		if err != nil {
			fmt.Println(err)
		}
	})
	baseMux.Handle("/", http.RedirectHandler("/custom", http.StatusMovedPermanently))
	server := http.Server{
		Addr:           ":8080",
		Handler:        baseMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 500,
		BaseContext:    f,
	}
	server.RegisterOnShutdown(func() { fmt.Println("server is shutting down") })
	return &server, errChan
}
