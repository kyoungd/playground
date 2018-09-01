package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// atomic is a class that is used for multithreading programming to gurantee that a variable is not corrupted between interrupt cycles.
// it syncs multitple threads to not collide for a same variable.
func main() {
	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops ", opsFinal)
}
