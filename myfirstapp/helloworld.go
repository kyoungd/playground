package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go World")
	oldSlc := []int{1, 3, 5, 7, 9, 11, 13}
	evenSlc := []int{2, 4, 6, 8}
	copy(evenSlc, oldSlc[2:])
	fmt.Println(evenSlc)
}
