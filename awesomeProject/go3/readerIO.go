package go3

import (
	"fmt"
	"io"
	"os"
)

func ReaderIo() {
	var i io.Reader
	_ = i
	/*	data, err := os.ReadFile("./awesomeProject/go3/text")
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		fmt.Println(string(data))*/

	f, err := os.Open("./awesomeProject/go3/text")
	defer f.Close()
	//os.MkdirAll()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	buf := make([]byte, 5)
	for {
		n1, err := f.Read(buf)
		fmt.Printf("%v, %v\n", n1, buf)
		if err != nil || n1 == 0 {
			fmt.Printf("%v\n", err)
			return
		}
	}

}
