package main

import (
	"fmt"
	"time"
)

func main() {

	request := make(chan int, 5)
	for ix := 1; ix <= 5; ix++ {
		request <- ix
	}

	limiter := time.NewTicker(time.Millisecond * 500)
	for req := range request {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}
}
