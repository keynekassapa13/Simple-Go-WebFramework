package routes

import (
	"fmt"
	"net/http"
	"encoding/json"

	// "goji.io"
	// "goji.io/pat"

	db "../db"
)


func GetUsers(res http.ResponseWriter, req *http.Request) {

	fmt.Println("[", req.Method, "] backend url", req.URL.Path)
	result := db.GetUsers()

	respBody, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	ResponseWithJSON(res, respBody, http.StatusOK)
}

func AddUser(res http.ResponseWriter, req *http.Request) {

	fmt.Println("[", req.Method, "] backend url", req.URL.Path)
	user := db.User{}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	result := db.AddUser(db.NewUser(user.Username, user.Password))

	if (result) {
		ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
	} else {
		ResponseWithJSON(res, []byte(`{"Result": "Error"}`), http.StatusOK)
	}
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("[", req.Method, "] backend url", req.URL.Path)

	user := db.User{}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	result := db.UpdateUser(user)

	if (result) {
		ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
	} else {
		ResponseWithJSON(res, []byte(`{"Result": "Error"}`), http.StatusOK)
	}
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("[", req.Method, "] backend url", req.URL.Path)

	user := db.User{}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	result := db.DeleteUser(user)

	if (result) {
		ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
	} else {
		ResponseWithJSON(res, []byte(`{"Result": "Error"}`), http.StatusOK)
	}
}
