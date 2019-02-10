package routes

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"

	"github.com/rivo/sessions"
	db "../db"
)

var s *sessions.Session

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
	s, _ = sessions.Start(res, req, true)

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
			Path:  "/",
			Expires: time.Now().Add(365 * 24 * time.Hour),
		})
		ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
	} else {
		ResponseWithJSON(res, []byte(`{"Result": "Error"}`), http.StatusOK)
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	if (s == nil) {
		s, _ = sessions.Start(res, req, true)
	}

	fmt.Println("[", req.Method, "] backend url", req.URL.Path)

	s.Delete("session_token")
	http.SetCookie(res, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:  	 "/",
		Expires: time.Now().Add(-365 * 24 * time.Hour),
	})

	ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
}

func IsLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	if (s == nil) {
		s, _ = sessions.Start(res, req, true)
	}
	fmt.Println("[", req.Method, "] isLoggedIn url", req.URL.Path)

	var tokenString interface {}

	tokenString = s.Get("session_token", tokenString)
	tokenString_c, err := req.Cookie("session_token")

	if (err != nil || tokenString == nil){
		fmt.Println(err)
		return false
	}

	if (tokenString == string(tokenString_c.Value)) {
		return true
	} else {
		return false
	}
}
