package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println("string 1: ", s)

	s, ok := i.(string)
	fmt.Println("string 2: ", s, ok)

	f, ok := i.(float64)
	fmt.Println("float 1: ", f, ok)

	f, _ = i.(float64)
	fmt.Println("float 2: ", f)

	f = i.(float64) // panic
	fmt.Println("float 3: ", f)
}
