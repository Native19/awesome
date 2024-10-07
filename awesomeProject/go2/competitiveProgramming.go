package go2

import "fmt"

type myStr struct {
	x, y int
}

func (*myStr) Get() (x, y int) {
	return x, y
}

func (myStr) Get2() (x, y int) {
	return x, y
}

func Main2() {
	bufferedChannels()

}

func Go2() {}

func bufferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
