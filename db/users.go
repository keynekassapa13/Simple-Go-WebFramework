package db

import (
  "fmt"
  "log"
  "time"

  // "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type User struct {
  Username string
  Password string
  CreatedAt time.Time
}

func NewUser(u string, p string) User {
    return User{
        Username: u,
        Password: p,
        CreatedAt: time.Now().Local(),
    }
}

func addUser(user User) User {
  fmt.Println("[ DB ]", "addUser ", user)
	c := session.DB("Auth").C("User")

  user.CreatedAt = time.Now().Local()

	err := c.Insert(user)
	if err != nil {
		log.Fatal(err)
	}

  return user
}

func GetUsers() []User {
  fmt.Println("[ DB ]", "getUsers ")
  c := session.DB("Auth").C("User")
  var users []User

  err := c.Find(bson.M{}).All(&users)
  if err != nil {
    log.Fatal(err)
  }

  return users
}
