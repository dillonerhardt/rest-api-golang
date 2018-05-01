package models

import (
	"log"
	"time"

	"github.com/dillonerhardt/rest-api-golang/db"
	"gopkg.in/mgo.v2/bson"
)

// User defines a user
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Mobile   string        `json:"mobile"`
	Gender   string        `json:"gender"`
	Password string        `json:"password"`
	Created  time.Time
}

const collection = "users"
const database = "a"

// GetUsers returns all users
func GetUsers() []User {
	s := db.DB.NewSession()
	defer s.Close()
	c := s.DB(database).C(collection)
	var users []User
	err := c.Find(bson.M{}).All(&users)
	if err != nil {
		log.Panic(err)
	}
	return users
}

// GetUser returns data for a user by id
func (u *User) GetUser() error {
	s := db.DB.NewSession()
	defer s.Close()
	c := s.DB(database).C(collection)
	return c.Find(bson.M{"_id": u.ID}).One(&u)
}

// GetUserByIdentifier returns data for a user by username
func (u *User) GetUserByIdentifier() error {
	s := db.DB.NewSession()
	defer s.Close()
	c := s.DB(database).C(collection)
	return c.Find(bson.M{"$or": []bson.M{
		bson.M{"username": u.Username},
		bson.M{"email": u.Email},
	}}).One(&u)
}

// CreateUser adds a new user to the database
func (u *User) CreateUser() error {
	s := db.DB.NewSession()
	defer s.Close()
	c := s.DB(database).C(collection)
	return c.Insert(u)
}
