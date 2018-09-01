package main

import (
	"fmt"
	"reflect"
)

type TMyFloat struct {
	X float64
	Y float64
	Z float64
}

func main() {
	var myFloat TMyFloat
	myFloat.X = 5.6
	myFloat.Y = 61.2
	myFloat.Z = 1.1
	inspectIfTypeMyFloat(myFloat)
}

func inspectIfTypeMyFloat(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println("type: ", v.Type()) // same as reflect.TypeOf(x1)
	fmt.Println("kind: ", v.Kind()) // same as reflect.TypeOf(x1)
	fmt.Println("Is type float64? ", v.Kind() == reflect.Float64, "  ", v.Kind())

	interfaceValue := v.Interface()
	switch t := interfaceValue.(type) {
	case TMyFloat:
		fmt.Println("TMyFloat : ", t)
	default:
		fmt.Println("default: ", t)
	}

	// x := interfaceValue.(type)
}
