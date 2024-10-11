package httpPractice

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func Serv() (*http.Server, chan error) {
	var errChan = make(chan error, 10)
	handlersErrCh := chan<- error(errChan)
	f := func(l net.Listener) context.Context {
		return context.WithValue(context.Background(), "errCh", handlersErrCh)
	}

	mux := http.NewServeMux()
	mux.Handle("/hand2", Handler2{})
	mux.HandleFunc("/hand1", Handler1{}.ServeHTTP)
	mux.HandleFunc("/id:5", Handler3{}.ServeHTTP)
	server := http.Server{
		Addr:           "127.0.0.1:8090",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 500,
		BaseContext:    f,
	}
	server.RegisterOnShutdown(func() { fmt.Println("server is shutting down") })
	return &server, errChan
}
