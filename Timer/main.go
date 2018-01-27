package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := tickCounter(ticker)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(3 * time.Second)
	fmt.Println("Exiting")
}

func tickCounter(ticker *time.Ticker) chan bool {
	done := make(chan bool)
	go func() {
		i := 0
	Loop:
		for {
			select {
			case t := <-ticker.C:
				i++
				fmt.Println("Count ", i, " at ", t)
			case <-done:
				break Loop
			}
			fmt.Println("Looping...")
		}
		fmt.Println("Exiting the tick counter")
	}()
	return done
}
