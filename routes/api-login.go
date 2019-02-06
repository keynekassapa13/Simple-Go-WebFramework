package routes

import (
	"fmt"
	"net/http"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("[", req.Method, "] backend url", req.URL.Path)
}
