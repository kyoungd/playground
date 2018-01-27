package main

import (
	"fmt"
	"young/FactoryPattern/appliances"
)

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
