package go2

import (
	"time"
)

func Chans() {
	//var ch2 chan int = make(chan int, 100000)
	/*	ch := make(chan int)
		var ch3 chan int
		close(ch)
		close(ch2)
		fmt.Println(len(ch3), cap(ch3), ch3)
		fmt.Println(len(ch), cap(ch), ch, <-ch)*/
	go func() {
		//var ch chan int
		/*		ch := make(chan int)
				close(ch)
				x := 10
				ch, ok <- x
				if ok {

				}
				select {
				case ch <- x:
					fmt.Println("write")
				default:
					fmt.Println("def")

				}*/

		//close(ch)
		//ch <- 1
		/*		x, y := <-ch
				if y {
					fmt.Println("success")
				}
				fmt.Println(ch == nil)
				//select {}
				x = <-ch
				fmt.Println(x)*/
		//fmt.Println("done")
		time.Sleep(1 * time.Hour)

	}()

	time.Sleep(time.Hour * 20)

	//go getChan(ch2, 1)
	//time.Sleep(20 * time.Second)
	/*	for {
		fmt.Println(<-ch2)
	}*/
	//fmt.Printf("ch2 is %T\n", ch2)
	//fmt.Println(cap(ch2))
	//<-ch
}

func getChan(ch chan int, num int) {
	for i := 0; ; i++ {
		ch <- i
		//fmt.Println(len(ch))
	}

}
