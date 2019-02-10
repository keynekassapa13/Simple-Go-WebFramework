package routes

import (
  "fmt"
  "os"
  "log"
  "net/http"
  "html/template"
  "path/filepath"

  // "github.com/rivo/sessions"
  "github.com/gorilla/mux"
)

var frontend_routes map[string]string
var auth_routes map[string]string

func FrontEndRoutes(r *mux.Router) {
  frontend_routes = map[string]string{
    "/"         : "index.html",
    "/index"    : "index.html",
    "/login"    : "auth/login.html",
    "/signup"   : "auth/signup.html",
    "/404"      : "error/404.html",
    "/500"      : "error/500.html",
  }
  auth_routes = map[string]string {
    "/auth/hi"       : "hello-world/index.html",
  }
  r.PathPrefix("/").HandlerFunc(serveTemplate)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
  fmt.Println("[", req.Method, "] frontend url", req.URL.Path)
  lp := filepath.Join("templates", "layout.html")
  var fp string

  if _, ok := frontend_routes[req.URL.Path]; ok {
    if (IsLoggedIn(res, req)) {
      fp = filepath.Join(
        "templates",
        filepath.Clean(
          auth_routes["/auth/hi"],
        ),
      )
    } else {
      fp = filepath.Join(
        "templates",
        filepath.Clean(
          frontend_routes[req.URL.Path],
        ),
      )
    }
  } else {
    if (IsLoggedIn(res, req)) {
      fp = filepath.Join(
        "templates",
        filepath.Clean(
          auth_routes[req.URL.Path],
        ),
      )
    } else {
      fp = filepath.Join(
        "templates",
        filepath.Clean(
          frontend_routes["/login"],
        ),
      )
    }
  }

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    fmt.Println(err.Error())
    if os.IsNotExist(err) {
      fp = filepath.Join(
        "templates",
        filepath.Clean(
          frontend_routes["/404"],
        ),
      )
    }
  }

  if info.IsDir() {
    fmt.Println("[", req.Method, "] frontend url", req.URL.Path, "is not found")
    fp = filepath.Join(
      "templates",
      filepath.Clean(
        frontend_routes["/404"],
      ),
    )
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    log.Println(err.Error())
    fp = filepath.Join(
      "templates",
      filepath.Clean(
        frontend_routes["/500"],
      ),
    )
  }

  if err := tmpl.ExecuteTemplate(res, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(res, http.StatusText(500), 500)
  }
}
