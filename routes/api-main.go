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

func RemoveTable(res http.ResponseWriter, req *http.Request) {
	fmt.Println("[", req.Method, "] backend url", req.URL.Path)

	type rmvTable struct { Dbname, Tname string }
	var p rmvTable

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&p)

	if err != nil {
		fmt.Println(err)
	}

	result := db.RemoveAll(p.Dbname, p.Tname)

	if (result) {
		ResponseWithJSON(res, []byte(`{"Result": "OK"}`), http.StatusOK)
	} else {
		ResponseWithJSON(res, []byte(`{"Result": "Error"}`), http.StatusOK)
	}
}
