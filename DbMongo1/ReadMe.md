Explains the mgo handler
https://godoc.org/labix.org/v2/mgo

Explains read in detail for mongodb.  Very useful.
https://docs.mongodb.com/manual/tutorial/query-documents/#read-operations-query-argument

	session, err := mgo.Dial("localhost:27017")
	// session, err := mgo.Dial("mongodb://127.0.0.1")
	defer session.Close()

	// reference collection
	personnel := session.DB("Hydra").C("Personnel")

	case "insert":
		newcr := CrewMember{ID: 10005, Name: "Keyla", SecurityClearance: 4, AccessCodes: []string{"CDF", "XYZ"}}
		err = personnel.Insert(newcr)
	case "update":
		err = personnel.Update(bson.M{"id": 10001}, bson.M{"$set": bson.M{"name": "Kenobi"}})
		PrintFatalError(err)
	case "delete":
		err = personnel.Remove(bson.M{"id": 10005})

	cm := CrewMember{}
	personnel.Find(bson.M{"id": 10001}).One(&cm)

	// Query with expression
	var crew Crew
	query := bson.M{ "clearanceLevel": bson.M{"$gt": 3,}, }
	err = personnel.Find(query).All(&crew)

	// select columns
	names := []struct { Name string `bson:"name"` }{}
	err = personnel.Find(query).Select(bson.M{"name": 1}).All(&names)
