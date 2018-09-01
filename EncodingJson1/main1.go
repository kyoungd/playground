package main

import (
	"encoding/json"
	"fmt"
)

type tCrewMember struct {
	ID                int
	Name              string
	SecurityClearance int
	AccessCodes       []string
}

type tShipInfo struct {
	ShipID    int
	ShipClass string
	Captain   tCrewMember
}

func main() {

	cm := tCrewMember{11, "Jaro", 10, []string{"abc123", "bcd234"}}
	si := tShipInfo{1, "Fighter", cm}

	bytesRep, err := json.Marshal(&si)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(string(bytesRep))

	// m := map[string]int{"item1": 1, "item2": 2}
	m := map[int]string{1: "item1", 2: "item2"}
	bm, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(string(bm))

	s := []tCrewMember{cm, tCrewMember{2, "Jim", 5, []string{"TLT", "RAR"}}}
	bSlice, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(string(bSlice))

	fmt.Println(" UNMARSHALL <<<<<<<<<<<<<<<<<< ")
	si0 := new(tShipInfo)
	json.Unmarshal(bytesRep, si0)
	fmt.Println(si0)

	m0 := make(map[int]string)
	json.Unmarshal(bm, &m0)
	fmt.Println(m0)

	b0 := []tCrewMember{}
	json.Unmarshal(bSlice, &b0)
	fmt.Printf("%#v", b0)
}
