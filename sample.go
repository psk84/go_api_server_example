package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	seq string
	id  string
	pwd string
}

func main() {
	session, err := mgo.Dial("daou-psk842.local:27017")
	if err != nil {
		panic(err)
		fmt.Println("error : ", err)
	}
	defer session.Close()

	fmt.Println("session", session)
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("user")
	err = c.Insert(&Person{"2", "realKang", "1234"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"seq": "1"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("id", result.id)
}
