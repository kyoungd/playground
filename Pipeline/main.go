package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
		fmt.Println("complete gen func")
	}()
	return out
}

func sq(in <-chan int) (<-chan int, <-chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	go func() {
		for n := range in {
			out <- n * n
			fmt.Println("-->", n*n)
		}
		close(out)
	}()
	return out, done
}

func main() {
	// Set up the pipeline.
	c := gen(2, 3, 4, 5)
	out, done := sq(c)

	// Consume the output.
	// fmt.Println(<-out)
	// fmt.Println(<-out)
	// fmt.Println(<-out)
	// fmt.Println(<-out)
	// fmt.Println(<-out)
	for x := range out {
		fmt.Println(x)
	}

	time.Sleep(1 * time.Second)
	go func() {
		<-done
		fmt.Println("Func Complete")
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("Program Complete")
	// for x := range out {
	// 	fmt.Println(x)
	// }
}
