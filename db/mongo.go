package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// MongoDB holds the database connection to MongoDB
type MongoDB struct {
	Session *mgo.Session
}

// DB holds the connection
var DB *MongoDB

// ConnectMongoDB connects a
func ConnectMongoDB(url string) {
	// create connection to MongoDB
	fmt.Println("Connecting to MongoDB")
	session, err := mgo.Dial(url)
	if err != nil {
		log.Panic(err)
	}
	DB = &MongoDB{
		Session: session,
	}
}

// NewSession returns a cloned session
func (db *MongoDB) NewSession() *mgo.Session {
	return db.Session.Clone()
}
