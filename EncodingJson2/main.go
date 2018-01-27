package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type CrewMember struct {
	ID                int
	Name              string
	SecurityClearance int
	AccessCodes       []string
}

type ShipInfo struct {
	ShipID    int
	ShipClass string
	Captain   CrewMember
}

func main() {
	f, err := os.Create("jfile.json")
	PrintFatalError(err)
	defer f.Close()

	cm := CrewMember{1000001, "Jaro", 10, []string{"ADA", "LAL"}}
	si := ShipInfo{1001, "Figther", cm}

	err = json.NewEncoder(f).Encode(si)
	PrintFatalError(err)

	ReadFile()
}

func ReadFile() {
	f, err := os.Open("jfile.json")
	PrintFatalError(err)
	defer f.Close()

	s0 := new(ShipInfo)
	err = json.NewDecoder(f).Decode(s0)
	PrintFatalError(err)
	fmt.Println(s0)
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happend. ", err)
	}
}
