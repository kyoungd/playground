package main

import (
	"fmt"
)

// Set of string and object type
type Set map[string]struct{}

func main() {
	s := make(Set)

	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	fmt.Println(getSetValue(s))
}

func getSetValue(s Set) string {
	var retVal string
	for k := range s {
		retVal = retVal + " " + k
	}
	return retVal
}
