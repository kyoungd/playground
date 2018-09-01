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

func ReflectSetInt(v interface{}) {

}

func main() {
	var myFloat TMyFloat
	myFloat.X = 5.6
	myFloat.Y = 61.2
	myFloat.Z = 1.1

	v := reflect.ValueOf(&myFloat)
	vpElem := v.Elem()
	if vpElem.Kind() == reflect.Struct {
		f := vpElem.FieldByName("X")
		if f.IsValid() {
			if f.CanSet() {
				if f.Kind() == reflect.Float64 {
					x := float64(0)
					if !f.OverflowFloat(x) {
						f.SetFloat(x)
					}
				}
			}
		}
	}
	fmt.Println(myFloat)

}
