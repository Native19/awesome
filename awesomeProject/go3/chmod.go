package go3

import (
	"fmt"
	"os"
)

func Chmod() {
	fmt.Println(0777)
	fmt.Println(777)
	fmt.Println(0b0101)

	file, err := os.Create("my.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.Chmod(0777)

}
