package go3

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Cont() {
	/*	ctxBack := context.Background()
		ctx, cfn := context.WithCancelCause(ctxBack)
		cfn(errors.New("reason"))
		//context.AfterFunc()

		//a := <-ctx.Value()

		//ctx.Done()

		cn := context.WithValue(ctxBack, "ctxBack", ctxBack)
		out := cn.Value("ctxBack")
		fmt.Println(out)
		context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
		_, cncl := context.WithCancel(context.Background())
		cncl()
		_, cncl2 := context.WithCancelCause(context.Background())
		cncl2(errors.New("cncl close"))
		//context.WithTimeout()
		ctx.Err()
		//context.Canceled
		//context.DeadlineExceeded
		//ctx.
		//cfn(errors.New("err"))
		ctx.Done()
		ctx.Deadline()*/
	//ctx.

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)

	<-ctx.Done()

	fmt.Println("ctx close")
}
