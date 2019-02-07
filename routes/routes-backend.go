package routes

import (
  // "fmt"

  "github.com/gorilla/mux"

)

func BackEndRoutes(r *mux.Router) {
  r.HandleFunc("/backend/hello", Hello)
  r.HandleFunc("/backend/getUsers", GetUsers)
}
