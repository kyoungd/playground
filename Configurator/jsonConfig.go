package HydrogenConfigurator

import (
	"encoding/json"
	"os"
)

type TJsonConfig struct{}

//
// DecodeJSONConfigFile (v: confirguration data, filename - configuration file name)
// it should read the config file and return the configuration as a json string with v.DecodeJSONConfigFile
func (t *TJsonConfig) DecodeJSONConfigFile(v interface{}, filename string) error {
	if filename == "" {
		filename = "configfile.json"
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	return json.NewDecoder(file).Decode(v)
}
