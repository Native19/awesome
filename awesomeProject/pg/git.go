package pg

import (
	"fmt"
	"os"
)

func Git() {
	for i := range 5 {
		fmt.Println(i)
	}
	os.Exit(0)
}
