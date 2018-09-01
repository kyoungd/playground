package main

import (
	"fmt"
	"playground/FactoryPattern/appliances"
)

// the stdin does not work with debugger.
// run it in the bash prompt screen.
// go run src/playground/FactoryPattern/main.go
func main() {
	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	var myType int
	fmt.Scan(&myType)
	// myType = 0

	myAppliance, err := appliances.CreateAppliance(myType)

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}
