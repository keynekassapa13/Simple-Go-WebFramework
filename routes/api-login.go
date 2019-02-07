package routes

import (
	"fmt"
	"net/http"
	"encoding/json"

	db "../db"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("[", req.Method, "] backend url", req.URL.Path)
}

func GetUsers(res http.ResponseWriter, req *http.Request) {

	fmt.Println("[", req.Method, "] backend url", req.URL.Path)
	result := db.GetUsers()

	respBody, err := json.MarshalIndent(result, "", "  ")

	if err != nil {
		fmt.Println(err)
	}

	ResponseWithJSON(res, respBody, http.StatusOK)
}
