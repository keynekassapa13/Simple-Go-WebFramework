package db

import (
  "fmt"
  // "log"
  // "reflect"

  "gopkg.in/mgo.v2"
)

var session, err = mgo.Dial("localhost")

func DeclareDb() {
  if err != nil {
          panic(err)
  }
  fmt.Println("[ DB ] Database is ready")

  session.SetMode(mgo.Monotonic, true)
}

func CloseDB() {
  session.Close()
}

func RemoveAll(
  dbname string,
  tname string,
) bool {
  fmt.Println("[ DB ]", "removeAll dbname=", dbname, " tname=", tname)
  c := session.DB(dbname).C(tname)

  err := c.DropCollection()
  if err != nil {
    fmt.Println("[ DB ] removeAll", err)
    return false
  }
  return true
}
