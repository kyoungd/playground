package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer func(ix int) {
				fmt.Println("Work end for ", ix)
				wg.Done()
			}(i)

			fmt.Println("Work start for ", i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	fmt.Println("Wait Complete")
}
