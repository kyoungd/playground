package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rand.Rand is a type Rand.  It is not a function.
// rand.New() returns a Rand type.
// Rand.Intn or other type of random numbe rgenerator can be used after that.
type customRand struct {
	*rand.Rand
	count int
}

// newCustomRand subclasses customRand
func newCustomRand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *customRand) RandRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *customRand) GetCount() int {
	return cr.count
}

func main() {
	cr := newCustomRand(time.Now().UnixNano())
	fmt.Println(cr.RandRange(5, 30))
	fmt.Println(cr.RandRange(5, 30))
	fmt.Println(cr.RandRange(5, 30))
	fmt.Println(cr.Intn(10))
	fmt.Println(cr.GetCount())
}
