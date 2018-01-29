package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go SayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}
	fmt.Println("Complete.")
}

func SayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i < n; i++ {
		c <- "Hello"
	}
	close(c)
}

// func main() {
// 	ch := make(chan string, 2)

// 	go waitAndReply(ch)
// 	ch <- "1"
// 	ch <- "2"
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)

// 	ch <- "3"
// 	ch <- "4"
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// func waitAndReply(ch chan string) {
// 	s := <-ch
// 	ch <- s
// }

// func main() {
// 	c := make(chan bool)
// 	go waitAndSay(c, "world")
// 	fmt.Println("Hello")
// 	// send a signal to c in order to allow waitAndSay to continue
// 	c <- true
// 	// wait to receive another sinal on c before we exit
// 	<-c
// }

// func waitAndSay(c chan bool, s string) {
// 	if b := <-c; b {
// 		fmt.Println(s)
// 	}
// 	c <- true
// }
