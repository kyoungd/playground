package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStruct struct {
		Field1 int     `alias:"f1" desc:"field number 1"`
		Field2 string  `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}

	mys := myStruct{2, "hello", 2.4}
	InspectStructType(mys)
	ChangeStructType(&mys)
	fmt.Println(mys)
}

func InspectStructType(i interface{}) {
	mysRValue := reflect.ValueOf(i)
	if mysRValue.Kind() != reflect.Struct {
		return
	}
	mysRType := reflect.TypeOf(i)

	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i)
		fieldRvalue := mysRValue.Field(i)
		fmt.Printf("%d. Name:'%s' Type: '%s' Value: '%v' ", i, fieldRType.Name, fieldRType.Type, fieldRvalue.Interface())
		fmt.Println(" ")
		fmt.Println("Struct tags, alias: ", fieldRType.Tag.Get("alias"), " description: ", fieldRType.Tag.Get("desc"))
	}
}

func ChangeStructType(i interface{}) {
	v := reflect.ValueOf(i)
	vpElem := v.Elem()
	if vpElem.Kind() == reflect.Struct {
		f := vpElem.FieldByName("Field1")
		if f.IsValid() {
			if f.CanSet() {
				f.SetInt(0)
				// if f.Kind() == reflect.Int64 {
				// 	x := int64(0)
				// 	if !f.OverflowInt(x) {
				// 		f.SetInt(x)
				// 	}
				// }
			}
		}
	}
}
