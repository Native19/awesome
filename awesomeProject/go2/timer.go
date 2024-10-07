package go2

import (
	"fmt"
	"time"
)

func Timer() {
	timer := time.NewTimer(5 * time.Second)
	ticker := time.NewTicker(time.Second)
	<-timer.C
	fmt.Println("timer is over")
	timer.Reset(5 * time.Second)
	<-timer.C
	timer.Stop()
	ticker.Stop()
}
