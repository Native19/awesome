package pg

import (
	"fmt"
	"os"
)

func Git() {
	mass := make([]int, 10)
	for i := range mass {
		fmt.Println(i)
	}
	os.Exit(0)
}
