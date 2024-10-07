package go2

import (
	"fmt"
	"time"
)

func Spinner1() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // Медленное вычисление
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	i := 0
	for {
		i++
		for range i {
			fmt.Printf("-")
		}
		fmt.Println()
		time.Sleep(delay)
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
