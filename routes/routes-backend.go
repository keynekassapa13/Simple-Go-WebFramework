package routes

import (
  // "fmt"

  "github.com/gorilla/mux"

)

func BackEndRoutes(r *mux.Router) {
  // api-main
  r.HandleFunc("/backend/hello", Hello)
  r.HandleFunc("/backend/removeTable", RemoveTable)

  // api-login
  r.HandleFunc("/backend/getUsers", GetUsers)
  r.HandleFunc("/backend/addUser", AddUser)
  r.HandleFunc("/backend/updateUser", UpdateUser)
  r.HandleFunc("/backend/deleteUser", DeleteUser)
  r.HandleFunc("/backend/login", Login)
  r.HandleFunc("/backend/isLoggedIn", IsLoggedIn)
}
