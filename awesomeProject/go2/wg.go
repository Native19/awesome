package go2

import (
	"fmt"
	"sync"
)

func Wg() {
	wgr := sync.WaitGroup{}
	wgr.Add(100)
	go func() {

		wgr.Add(-100)
		//defer wgr.Done()

	}()
	fmt.Println("11111")
	wgr.Wait()

	//wgr.Done()
}
