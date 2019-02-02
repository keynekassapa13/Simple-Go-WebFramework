package routes

import (
  "fmt"
  "os"
  "log"
  "net/http"
  "html/template"
  "path/filepath"

  "github.com/gorilla/mux"
)

func FrontEndRoutes(r *mux.Router) {
  r.PathPrefix("/").HandlerFunc(serveTemplate)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
  fmt.Println("[", req.Method, "] frontend url", req.URL.Path)
  lp := filepath.Join("templates", "layout.html")
  var fp string
  
  if req.URL.Path == "/" {
    fp = filepath.Join("templates", filepath.Clean("index.html"))
  } else {
    fp = filepath.Join("templates", filepath.Clean(req.URL.Path + ".html"))
  }

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    fmt.Println(err.Error())
    if os.IsNotExist(err) {
      http.NotFound(res, req)
      return
    }
  }

  if info.IsDir() {
    fmt.Println(info)
    http.NotFound(res, req)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    log.Println(err.Error())
    http.Error(res, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(res, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(res, http.StatusText(500), 500)
  }
}
