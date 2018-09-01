package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := a[2:4]
	fmt.Println("length ", len(s))
	fmt.Println("capacity ", cap(s))
	fmt.Println("s ", s)
	s = s[2:4]
	fmt.Println("length ", len(s))
	fmt.Println("capacity ", cap(s))
	fmt.Println("s ", s)
	s = s[2:4]
	fmt.Println("length ", len(s))
	fmt.Println("capacity ", cap(s))
	fmt.Println("s ", s)
	b := append(a, s...)
	fmt.Println("length ", len(b))
	fmt.Println("capacity ", cap(b))
	fmt.Println("s ", b)
	c := make([]int, len(b), cap(b))
	copy(c, b)
	fmt.Println("length ", len(c))
	fmt.Println("capacity ", cap(c))
	fmt.Println("s ", c)
}
