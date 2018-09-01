// Use the built-in synchronization features of  goroutines and
// channels to synchronize access to shared state
// across multiple goroutines. This channel-based
// approach aligns with Go's ideas of sharing memory by
// communicating and having each piece of data owned
// by exactly 1 goroutine.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

// In this example our state will be owned by a single goroutine. This will guarantee that the data is never
// corrupted with concurrent access. In order to read or write that state, other goroutines will send messages
// to the owning goroutine and receive corresponding replies. These `readOp` and `writeOp` `struct`s
// encapsulate those requests and a way for the owning goroutine to respond.
func main() {

	// As before we'll count how many operations we perform.
	var readOpCount uint64
	var writeOpCount uint64

	// The `reads` and `writes` channels will be used by
	// other goroutines to issue read and write requests,
	// respectively.
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// Here is the goroutine that owns the `state`, which is a map as in the previous example but now private
	// to the stateful goroutine. This goroutine repeatedly selects on the `reads` and `writes` channels,
	// responding to requests as they arrive. A response is executed by first performing the requested
	// operation and then sending a value on the response channel `resp` to indicate success (and the desired
	// value in the case of `reads`).
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				if value, ok := state[read.key]; ok {
					read.resp <- value
				} else {
					read.resp <- 0
				}
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	// This starts 100 goroutines to issue reads to the state-owning goroutine via the `reads` channel.
	// Each read requires constructing a `readOp`, sending it over the `reads` channel, and the receiving the
	// result over the provided `resp` channel.
	for ix := 1; ix <= 100; ix++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOpCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We start 10 writes as well, using a similar
	// approach.
	for ix := 1; ix <= 100; ix++ {
		go func() {
			for {
				write := &writeOp{
					key:   rand.Intn(5),
					value: ix,
					resp:  make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOpCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second.
	time.Sleep(time.Second * 2)

	// Finally, capture and report the op counts.
	readOpCountFinal := atomic.LoadUint64(&readOpCount)
	writeOpCountFinal := atomic.LoadUint64(&writeOpCount)
	fmt.Println("readOpCount:", readOpCountFinal)
	fmt.Println("writeOpCount:", writeOpCountFinal)
}
