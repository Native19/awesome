package testMiddleware

import (
	"fmt"
	"time"

	"awesome/awesomeProject/httpPractice"
)

func TestMiddleware() {
	t1 := time.Now()
	fmt.Println("Server starting")
	serv, erCh := httpPractice.Serv()
	serv.Handler = newRouter()
	go func(erCh chan error) {
		for err := range erCh {
			fmt.Println(err.Error())
		}
	}(erCh)

	go func() {
		if err := serv.ListenAndServe(); err != nil {
			return
		}
	}()
	fmt.Println("Server started")

	time.Sleep(100 * time.Second)
	//stopServ, _ := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Printf("время работы сервера: %v\n", time.Since(t1).String())
	//_ = serv.Shutdown(stopServ)
	fmt.Printf("время работы сервера и выключения: %v\n", time.Since(t1).String())
}
