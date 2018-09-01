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
	Name1, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(name, server string) (int, error) {
	// simulate searching
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	if v, ok := scMapping[name]; !ok {
		// panic ("Acckkkkkkkkk!!!")
		return -1, findError{name, server, "Crew member not found. "}
	} else {
		return v, nil
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// defer function() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("A panic recoverd", err)
	// 	}
	// }()
	if clearance, err := findSC("Ruko", "Server 1"); err != nil {
		fmt.Println("Error occured while searching ", err)
		msg := err.Error()
		fmt.Println(msg)
		if v, ok := err.(findError); ok {
			fmt.Println("Server name ", v.Server)
			fmt.Println("Crew mebmer name ", v.Name1)
		} else {
			fmt.Println("Clearance level found: ", clearance, " Error code ", err)
		}
	}
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("A Panic recovered ", err)
	// 	}
	// }()

	// c1 := make(chan int)
	// c2 := make(chan int)

	// name := "James"

	// fmt.Println("Searching for James...")

	// go findSC(name, "Server 1", c1)
	// go findSC(name, "Server 2", c2)

	// select {
	// case sc := <-c1:
	// 	fmt.Println(name, " has a security clearance of ", sc, " found in server 1")
	// case sc := <-c2:
	// 	fmt.Println(name, " has a security clearance of ", sc, " found in server 2")
	// default:
	// 	fmt.Println("default called.")
	// 	// case <-time.After(25 * time.Second):
	// 	// 	fmt.Println("Search time out!")
	// }
	// fmt.Println("Keep going.")
	// time.Sleep(60 * time.Second)
	// fmt.Println("Done.")
}
