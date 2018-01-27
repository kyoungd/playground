package main

import (
	"encoding/xml"
	"fmt"
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	cm := CrewMember{1000001, "Jaro", 10, []string{"ADA", "LAL"}}
	si := ShipInfo{ShipID: 1001, ShipClass: "Figther", Captain: cm}

	// b, err := xml.Marshal(&si)
	b, err := xml.MarshalIndent(&si, " ", "	")
	PrintFatalError(err)
	fmt.Println(xml.Header, string(b))

	s := []int{1, 2, 3, 4}
	ix, err := xml.Marshal(&s)
	PrintFatalError(err)
	fmt.Println(xml.Header, string(ix))

	si0 := new(ShipInfo)
	err2 := xml.Unmarshal(b, &si0)
	PrintFatalError(err2)
	fmt.Println(si0)

	ix0 := make([]int, 0)
	err3 := xml.Unmarshal(ix, &ix0)
	PrintFatalError(err3)
	fmt.Printf("%#v", ix0)

}

func PrintFatalError(err error) {
	if err != nil {
		panic(fmt.Sprintf("PANIC: Error happend. %", err))
	}
}
