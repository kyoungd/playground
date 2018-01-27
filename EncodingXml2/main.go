package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type CrewMember struct {
	ID                int      `xml:"id,omitempty"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearanceLevel"`
	AccessCodes       []string `xml:"accesscodes>code"`
}

type ShipInfo struct {
	XMLName   xml.Name `xml:"Root"`
	ShipID    int
	ShipClass string
	Captain   CrewMember
}

func main() {
	f, err := os.Create("jfile.xml")
	PrintFatalError(err)
	defer f.Close()

	cm := CrewMember{1000001, "Jaro", 10, []string{"ADA", "LAL"}}
	si := ShipInfo{ShipID: 1001, ShipClass: "Figther", Captain: cm}

	err = xml.NewEncoder(f).Encode(si)
	PrintFatalError(err)

	ReadFile()
}

func ReadFile() {
	f, err := os.Open("jfile.xml")
	PrintFatalError(err)
	defer f.Close()

	s0 := new(ShipInfo)
	err = xml.NewDecoder(f).Decode(s0)
	PrintFatalError(err)
	fmt.Println(s0)
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happend. ", err)
	}
}
