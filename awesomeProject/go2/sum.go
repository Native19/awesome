package go2

import (
	"fmt"
	"sync"
	"time"
)

const length = 200000
const goroutines = 100

var slice []int = make([]int, length*goroutines)

func Sum() {
	/*	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()*/

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f222222", r)
			}
		}()
		time.Sleep(1 * time.Second)
	}()

	go func() {
		defer fmt.Println("-----------")
		//os.Exit(0)
		for ty := range 5 {
			g := 5 / ty
			fmt.Println(g)
		}

	}()

	time.Sleep(20 * time.Second)
	//mu := sync.Mutex{}
	//mu.Lock()
	//mu.Unlock()
	//mu.Unlock()

	/*rwmu := sync.RWMutex{}
	//rwmu.Unlock()
	rwmu.RLock()
	//rwmu.RLock()
	fmt.Println("RUnlock")
	rwmu.RUnlock()
	rwmu.Lock()
	//rwmu.Lock()
	fmt.Println("lock")
	rwmu.Unlock()*/
	start := time.Now()
	type chanSave struct {
		ch      chan int
		isClose bool
	}
	forSum := make(chan int, goroutines)
	chS := chanSave{forSum, false}
	wg := sync.WaitGroup{}
	for i := range goroutines {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			for m := range length {
				slice[num*length+m] = m
			}
		}(i)
	}
	wg.Wait()
	go func(save *chanSave) {
		//rnd := rand.Intn(10) + 1
		//time.Sleep(time.Millisecond * time.Duration(rnd))
		//fmt.Println(time.Since(start), " close")

		wg.Wait()
		chS.isClose = true
		close(forSum)
	}(&chS)

	for gi := range goroutines {
		wg.Add(1)
		go func(num int, save *chanSave) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()
			defer wg.Done()
			sum := 0
			for i := range length {
				sum += slice[num*length+i]
			}
			if !save.isClose {
				save.ch <- sum
			}
		}(gi, &chS)
	}

	fullSum := 0
	for v := range chS.ch {
		fullSum += v
	}
	stop := time.Since(start)

	fmt.Println(fullSum)
	fmt.Println(stop, " stop")
}
