package db

import (
  "fmt"
  "log"
  "time"
  "github.com/rs/xid"

  // "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  Id xid.ID `json:"_id", string`
  Username string `json:"username"`
  Password string `json:"password"`
  CreatedAt time.Time `json:"createdat"`
}

func NewUser(u string, p string) User {
  h, _ := HashPassword(p)
  return User{
    Id: xid.New(),
    Username: u,
    Password: h,
    CreatedAt: time.Now().Local(),
  }
}

func AddUser(user User) bool {
  fmt.Println("[ DB ]", "addUser ", user)
	c := session.DB("Auth").C("User")

  user.CreatedAt = time.Now().Local()

	err := c.Insert(user)
	if err != nil {
		log.Fatal(err)
    return false
	}

  return true
}

func GetUsers() []User {
  fmt.Println("[ DB ]", "getUsers ")
  c := session.DB("Auth").C("User")
  var users []User

  err := c.Find(bson.M{}).Sort("createdat").Select(bson.M{"username": 1, "_id": 1, "createdat": 1}).All(&users)
  if err != nil {
    log.Fatal(err)
  }

  return users
}

func UpdateUser(user User) bool {
  fmt.Println("[ DB ]", "UpdateUser ", user)
  c := session.DB("Auth").C("User")

  err := c.Update(
    bson.M{"username": user.Username},
    bson.M{"$set": bson.M{
      "username": user.Username,
      "password": user.Password,
    }},
  )
  if err != nil {
		fmt.Println(err)
    return false
	}

  return true
}

func DeleteUser(user User) bool {
  fmt.Println("[ DB ]", "DeleteUser ", user)
  c := session.DB("Auth").C("User")

  err := c.Remove(
    bson.M{"username": user.Username},
  )

  if err != nil {
		fmt.Println(err)
    return false
	}

  return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
