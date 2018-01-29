package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("Wait 2 second")

	<-timer1.C
	fmt.Println("Waited 2 second")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 2 second expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("time2 has stopped")
	}
}
