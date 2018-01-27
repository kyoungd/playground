package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = append(a[:2], a[4:]...)
	fmt.Println(a)

	// a := []int{1, 2, 3, 4, 5, 6}
	// b := a[2:4]
	// c := a[:3]
	// d := a[3:]
	// s := make([]int, 3)
	// copy(s, d[:3])
	// s[0] = 10
	// fmt.Println("Slices a:", a, " b:", b, " c:", c, " d:", d)
	// fmt.Println("Slices a:", cap(a), " b:", cap(b), " c:", cap(c), " d:", cap(d))
}
