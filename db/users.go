package db

import (
  "fmt"
  "log"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type User struct {
  Name string
  Phone string
}

func addUser(session *mgo.Session, user *User) {
  fmt.Println("[ DB ]", "addUser ", user)
	c := session.DB("Auth").C("User")

	err := c.Insert(user)
	if err != nil {
		log.Fatal(err)
	}

}


func getUsers(session *mgo.Session) {
  fmt.Println("[ DB ]", "getUsers ")
  c := session.DB("Auth").C("User")
  var users []User

  err := c.Find(bson.M{}).All(&users)
  if err != nil {
    log.Fatal(err)
  }

  return users
}
