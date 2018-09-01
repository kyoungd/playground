package main

import "fmt"

// This is a nice collection that shows how GO handles channel, and a proper way to close channel.
func main() {
	job := make(chan int, 3)
	done := make(chan bool)

	go func() {
		for {
			// item and more is waiting for an input 'job'.
			item, more := <-job
			if more == true {
				fmt.Println("there is more. ", item)
			} else {
				fmt.Println("no more")
				// assigned 'done'
				done <- true
				return
			}
		}
	}()

	for ix := 0; ix < 3; ix++ {
		// 'job' is assigned here.
		job <- ix
		fmt.Println("times: ", ix)
	}
	// close channel on 'job'.  item, more := <-job (nil, true)
	close(job)
	fmt.Println("complete sent")
	// waiting for the 'done' to be assigned
	<-done

}
