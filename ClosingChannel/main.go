package main

import "fmt"

func main() {
	job := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for {
			item, more := <-job
			if more == true {
				fmt.Println("there is more. ", item)
			} else {
				fmt.Println("no more")
				done <- true
				return
			}
		}
	}()

	for ix := 0; ix < 3; ix++ {
		job <- ix
		fmt.Println("times: ", ix)
	}
	close(job)
	fmt.Println("complete sent")
	<-done

}
