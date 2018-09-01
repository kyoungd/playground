package main

import (
	"encoding/xml"
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CrewMember struct {
	ID                int      `bson:"id,omitempty"`
	Name              string   `bson:"name"`
	SecurityClearance int      `bson:"clearanceLevel"`
	AccessCodes       []string `bson:"accessCodes"`
}

type Crew []CrewMember

type ShipInfo struct {
	XMLName   xml.Name `xml:"Root"`
	ShipID    int
	ShipClass string
	Captain   CrewMember
}

func crewMemberAction(personnel *mgo.Collection, action string) {
	var err error

	switch action {
	case "insert":
		// insert
		newcr := CrewMember{ID: 10005, Name: "Keyla", SecurityClearance: 4, AccessCodes: []string{"CDF", "XYZ"}}
		err = personnel.Insert(newcr)
		PrintFatalError(err)
	case "update":
		// update
		err = personnel.Update(bson.M{"id": 10001}, bson.M{"$set": bson.M{"name": "Kenobi"}})
		PrintFatalError(err)
	case "delete":
		// delete
		err = personnel.Remove(bson.M{"id": 10005})
		PrintFatalError(err)
	default:
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	session, err := mgo.Dial("localhost:27017")
	// session, err := mgo.Dial("mongodb://127.0.0.1")
	PrintFatalError(err)
	defer session.Close()

	// reference collection
	personnel := session.DB("Hydra").C("Personnel")

	crewMemberAction(personnel, "") // do nothing

	// GEt number of documents in the collection
	n, _ := personnel.Count()
	fmt.Println("Number of personnel is ", n)

	// Perform simple query
	cm := CrewMember{}
	personnel.Find(bson.M{"id": 10001}).One(&cm)
	fmt.Println(cm)

	// Query with expression
	query := bson.M{
		"clearanceLevel": bson.M{"$gt": 3},
	}

	var crew Crew
	err = personnel.Find(query).All(&crew)
	PrintFatalError(err)
	fmt.Println(crew)

	// select columns
	names := []struct {
		Name string `bson:"name"`
	}{}
	err = personnel.Find(query).Select(bson.M{"name": 1}).All(&names)
	PrintFatalError(err)
	fmt.Println(names)

	var wg sync.WaitGroup
	count, _ := personnel.Count()
	wg.Add(count)
	for i := 10001; i <= 10000+count; i++ {
		go readID(i, session.Copy(), &wg)
	}
	wg.Wait()
}

func readID(id int, ses *mgo.Session, wg *sync.WaitGroup) {
	defer func() {
		ses.Close()
		wg.Done()
	}()
	p := ses.DB("Hydra").C("Personnel")
	cm := CrewMember{}
	err := p.Find(bson.M{"id": id}).One(&cm)
	PrintFatalError(err)
	fmt.Println(cm)
}

func PrintFatalError(err error) {
	if err != nil {
		msg := fmt.Sprintf("Error happend. %s", err)
		panic(msg)
	}
}
