package main

import (
	"fmt"
	"math/rand"
	"time"
)

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

type findError struct {
	Name, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(name, server string) (int, error) {
 	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	if v, ok := scMapping[name]; !ok {
		return -1, findError{name, server, "Crew member not found"}
	} else {
		return v, nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	defer func() {
		if err != recover(); err != {
			fmt.Println("A panic recovered ", err)
		}
	}
	clearance, err := find("Ruko", "Server 1")
	fmt.Println("Clearance level found: ", clearance, " error code: ", err)
}

// func findSC(name, server string, c chan int) {
// 	// simulate searching
// 	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
// 	// return security clearance from map
// 	c <- scMapping[name]
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	c1 := make(chan int)
// 	c2 := make(chan int)

// 	name := "James"

// 	fmt.Println("Searching for James...")

// 	go findSC(name, "Server 1", c1)
// 	go findSC(name, "Server 2", c2)

// 	select {
// 	case sc := <-c1:
// 		fmt.Println(name, " has a security clearance of ", sc, " found in server 1")
// 	case sc := <-c2:
// 		fmt.Println(name, " has a security clearance of ", sc, " found in server 2")
// 		// default:
// 		// 	fmt.Println("default called.")
// 		// case <-time.After(25 * time.Second):
// 		// 	fmt.Println("Search time out!")
// 	}
// 	fmt.Println("Keep going.")
// 	time.Sleep(6 * time.Second)
// 	fmt.Println("Done.")
// }
