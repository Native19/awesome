package go2

import (
	"fmt"
	"time"
)

func Pract() {

	defer time.Sleep(1000 * time.Hour)
	fmt.Println("Hello World")
	panic("panic")
	/*	var m = make(map[int]int)
		ch := make(chan int, 1)
		wg := &sync.WaitGroup{}

		go func(m2 map[int]int, ch2 chan int, wg2 *sync.WaitGroup) {
			wg.Add(1)
			fmt.Println("goroutine start")
			ch <- 1
			fmt.Println("goroutine start writing")
			for i := range 1000 {
				m[i] = i * 100
			}
			fmt.Println("goroutine stop writing")
			<-ch
			wg.Done()
		}(m, ch, wg)

		ch <- 1
		fmt.Println("main start writing")
		for k := range 1000 {
			m[k] = k * 10000
		}
		fmt.Println("main stop writing")
		<-ch
		wg.Wait()

		close(ch)
		fmt.Println(len(ch), cap(ch))*/

	//y := 100
	//ch := make(chan int)

	//chr := make(<-chan int, 10)
	//ch := make(chan int)

	//ch2 := chan int(chr)
	//var ch3 <-chan int = ch
	//close(ch)
	//close(ch)
	//<-ch3
	/*		for i := 0; i < 10; i++ {
			ch <- i
		}*/
	//close(ch)

	/*	for i := range ch {
		fmt.Println(i)
	}*/

	/*	go func() {
		time.Sleep(24 * time.Second)
		ch <- 5
	}()*/
	//time.Sleep(3 * time.Second)
	//fmt.Println(<-ch)
	/*	for i := 0; i < 10; i++ {
			b, ok := <-ch
			fmt.Println(b, ok)
		}
		b, ok := <-ch
		fmt.Println(b, ok)*/

	/*	fmt.Println("Hello practice")
		defer fmt.Println("defer practice")
		go func() {
			defer fmt.Println("defer goroutine")
			panic("goroutine panic")
			os.Exit(0)
		}()
		time.Sleep(5 * time.Second)*/
	//var ch chan int
	/*		go func(ch2 *chan int) {
				time.Sleep(2 * time.Second)
				*ch2 = make(chan int)
				*ch2 <- 5
				fmt.Println("goroutine finished")
			}(&ch)

			go func(ch2 chan int) {
				time.Sleep(2 * time.Second)
				ch2 = make(chan int)
				ch2 <- 5
				fmt.Println("goroutine finished")
			}(ch)*/

	/*	fmt.Println("start waiting")
		Loop:
			for {
				select {
				case <-ch:
					fmt.Println("ch unlock")
					break Loop
				case <-time.After(1 * time.Second):
					fmt.Println("timer")
				default:
					fmt.Println("default")

				}
			}

			a := <-ch
			_ = a
			fmt.Println("main finished")*/

	/*	var ch chan int = make(chan int, 1)
			go func(ch2 *chan int) {
				time.Sleep(2 * time.Second)
				close(*ch2)
				fmt.Println("goroutine is over")
			}(&ch)

			fmt.Println("main is waiting")
		Loop:
			for {
				select {
				case a, ok := <-ch:
					fmt.Println(a, ok)
					break Loop
				case ch <- 5:
					fmt.Println("write in chan")
				}
			}
			fmt.Println("main is done")*/
}
