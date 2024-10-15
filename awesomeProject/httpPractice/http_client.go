package httpPractice

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Ht() {
	client := http.Client{}
	/*	req := http.Request{}
		res, err := client.Do(&req)*/

	resp, err := client.Get("https://golangforall.com/en/")
	if err != nil {
		log.Fatal(err)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}
