package main

import (
	"fmt"
	"time"
)

func main() {
	ticker1 := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker1.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Program Exits")

}
