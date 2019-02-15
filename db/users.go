package db

import (
  "fmt"
  "log"
  "time"

  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/dgrijalva/jwt-go"
  "github.com/rs/xid"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  Id xid.ID `json:"_id", string`
  Username string `json:"username"`
  Password string `json:"password"`
  CreatedAt time.Time `json:"createdat"`
}

type JwtToken struct {
    Token string `json:"token"`
}

func NewUser(u string, p string) User {
  h, err := HashPassword(p)
  if err != nil {
      fmt.Println(err)
  }
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

  index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

  err = c.EnsureIndex(index)
	if err != nil {
		fmt.Println(err)
    return false
	}

  user.CreatedAt = time.Now().Local()

	err := c.Insert(user)
	if err != nil {
		fmt.Println(err)
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

  index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

  err = c.EnsureIndex(index)
	if err != nil {
		fmt.Println(err)
    return false
	}

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

func findUserByUsername(username string) User {
  fmt.Println("[ DB ]", "findUsers ", username)
  c := session.DB("Auth").C("User")

  var user User
  err = c.Find(bson.M{"username": username}).One(&user)

  return user
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(user User) string {
  user_auth := findUserByUsername(user.Username)

  if (CheckPasswordHash(user.Password, user_auth.Password)) {
    fmt.Println("CheckPasswordHash yes")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "username": user.Username,
      "password": user.Password,
    })

    tokenString, err := token.SignedString([]byte("secret"))

    if err != nil {
        fmt.Println(err)
    }
    return tokenString
  }
  return ""
}
