package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	UserId string        `bson:"userId"`
	Pwd    string        `bson:"pwd"`
}

func insertUser(name string, pwd string) bool {
	fmt.Println("Insert User : ", name)
	var result bool = false

	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
		fmt.Println("error : ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	err = session.DB("test").DropDatabase()
	if err != nil {
		panic(err)
	}

	c := session.DB("test").C("user")
	// Index
	index := mgo.Index{
		Key:        []string{"userId", "pwd"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	value := &User{Id: bson.NewObjectId(), UserId: name, Pwd: pwd}
	fmt.Println("value : ", value)

	// Insert Datas
	err = c.Insert(value)

	if err != nil {
		result = false
	} else {
		result = true
	}

	fmt.Println("Insert result : ", result)

	return result
}

func selectUser(name string, pwd string) *User {
	fmt.Println("Select User : ", name)
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
		fmt.Println("error : ", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("user")

	// Query
	var result User
	err = c.Find(bson.M{"userId": name}).One(&result)

	if err != nil {
		return nil
	}
	fmt.Println("Results All: ", result)

	return &result
}
