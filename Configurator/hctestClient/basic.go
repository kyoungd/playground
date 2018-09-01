package main

import (
	"fmt"
	"playground/Configurator"
)

type Confs struct {
	TS      string  `name:"testString" json:"testString" xml:"testString"`
	TB      bool    `name:"testBool" json:"testBool" xml:"testBool"`
	TF      float64 `name:"testFloat" json:"testFloat" xml:"testFloat"`
	TestInt int
}

type IConfigJson interface {
	DecodeJSONConfigFile(v interface{}, filename string) error
}

func loadFile(loadOne IConfigJson, v interface{}, filename string) error {
	return loadOne.DecodeJSONConfigFile(v, filename)
}

func main() {
	configstruct := new(Confs)
	loadFile(&HydrogenConfigurator.TJsonConfig{}, configstruct, "configfile.json")
	// HydrogenConfigurator.DecodeJSONConfigFile(configstruct, "configfile.json")

	fmt.Println(*configstruct)
	if configstruct.TB {
		fmt.Println("bool is true")
	}
	fmt.Println(float64(4.8 * configstruct.TF))
	fmt.Println(5 * configstruct.TestInt)
	fmt.Println(configstruct.TS)
}
