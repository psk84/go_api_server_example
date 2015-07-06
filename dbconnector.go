package main

import (
	"gopkg.in/mgo.v2"
)

func dbConnection() (*Session, error) {
	session, err := mgo.Dial("daou-psk842.local:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return session, err
}
