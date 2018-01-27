// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for jb := range jobs {
		fmt.Println("Start ID %d Job %d ", id, jb)
		time.Sleep(time.Second * 2)
		fmt.Println("Complete ID %d Job %d ", id, jb)
		results <- jb * 10
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for ix := 1; ix <= 3; ix++ {
		go worker(ix, jobs, results)
	}

	for ix := 0; ix < 5; ix++ {
		jobs <- (ix + 1)
	}
	close(jobs)

	go func() {
		for r := range results {
			fmt.Println("RESULT: ", r)
		}
	}()

	time.Sleep(time.Second * 30)
	close(results)
}
