package main

import (
	"fmt"
	"reflect"
)

type YourT1 struct{}

func (y YourT1) MethodBar() {
	fmt.Println("MethodBar()")
}

type YourT2 struct{}

func (y YourT2) MethodFoo(i int, oo string) {
	fmt.Printf("MethodFoo() %d %s", i, oo)
	fmt.Println("")
}

func Invoke(any interface{}, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

func AvailableMethod(any interface{}) {
	mysRType := reflect.TypeOf(any)
	for i := 0; i < mysRType.NumMethod(); i++ {
		m := mysRType.Method(i)
		fmt.Printf("%+v", m)
		fmt.Println("")
	}
}

func main() {
	var yt2 YourT2
	AvailableMethod(yt2)
	Invoke(YourT2{}, "MethodFoo", 10, "abc")
	Invoke(YourT1{}, "MethodBar")
}
