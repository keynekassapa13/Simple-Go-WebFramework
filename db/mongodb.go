package db

import (
  "fmt"

  "gopkg.in/mgo.v2"
  // "gopkg.in/mgo.v2/bson"
)

func DeclareDb() {

  session, err := mgo.Dial("localhost")
  if err != nil {
          panic(err)
  }
  fmt.Println("[ DB ] Database is ready")
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)

  addUser(session, &User{"Ale", "+55 53 8116 9639"})
  addUser(session, &User{"Cla", "+55 53 8402 8510"})
  getUsers(session)
}
