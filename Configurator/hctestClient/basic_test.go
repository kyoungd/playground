package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TJsonConfig struct{}

func (t *TJsonConfig) DecodeJSONConfigFile(i interface{}, filename string) error {
	v := reflect.ValueOf(i)
	vpElem := v.Elem()
	if vpElem.Kind() == reflect.Struct {
		testInt := vpElem.FieldByName("TestInt")
		if testInt.IsValid() && testInt.CanSet() {
			testInt.SetInt(2)
		}
		ts := vpElem.FieldByName("TS")
		if ts.IsValid() && ts.CanSet() {
			ts.SetString("Hello there")
		}
		tb := vpElem.FieldByName("TB")
		if tb.IsValid() && tb.CanSet() {
			tb.SetBool(false)
		}
		tf := vpElem.FieldByName("TF")
		if tf.IsValid() && tf.CanSet() {
			tf.SetFloat(3.14159265)
		}
	}
	return nil
}

func TestLoadConfigFile(t *testing.T) {
	configstruct := new(Confs)
	loadFile(&TJsonConfig{}, configstruct, "configfile.json")
	// HydrogenConfigurator.DecodeJSONConfigFile(configstruct, "configfile.json")

	assert.Equal(t, configstruct.TestInt, 3, "integer value is not equal to zero")
	assert.Equal(t, configstruct.TF, 3.14159265, "float point value must be pi")

	// fmt.Println(*configstruct)
	// if configstruct.TB {
	// 	fmt.Println("bool is true")
	// }
	// fmt.Println(float64(4.8 * configstruct.TF))
	// fmt.Println(5 * configstruct.TestInt)
	// fmt.Println(configstruct.TS)
}
