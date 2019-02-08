package routes

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"

	"github.com/rivo/sessions"
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

func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("[", req.Method, "] backend url", req.URL.Path)
	s, _ := sessions.Start(res, req, true)

	user := db.User{}

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	result := db.Login(user)

	if (result != "") {
		s.Set("session_token", result)
		http.SetCookie(res, &http.Cookie{
			Name:    "session_token",
			Value:   result,
			Expires: time.Now().Add(12000 * time.Second),
		})
		ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
	} else {
		ResponseWithJSON(res, []byte(`{"Result": "Error"}`), http.StatusOK)
	}
}

func IsLoggedIn(res http.ResponseWriter, req *http.Request) {
	s, _ := sessions.Start(res, req, false)
	tokenString := s.Get("session_token", nil)
	tokenString_c, _ := req.Cookie("session_token")

	if (tokenString == string(tokenString_c.Value)) {
		ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
	} else {
		ResponseWithJSON(res, []byte(`{"Result": "Error"}`), http.StatusOK)
	}
}
