package main

import(
  "fmt"
  "log"
  "time"
  "flag"
  "net/http"

  "github.com/gorilla/mux"

  routes "./routes"
)

func main() {

  fmt.Println("Starting the server...")

  r := mux.NewRouter()

  declareRoutes(r)
  http.Handle("/", r)

  fmt.Println("[ SERVER ] Server is ready at port 9000")
  srv := &http.Server{
        Handler:      r,
        Addr:         ":9000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
  log.Fatal(srv.ListenAndServe())
}

func declareRoutes(r *mux.Router) {
  var dir string
  flag.StringVar(&dir, "static", "./static/", "the directory to serve files from. Defaults to the current dir")
  flag.Parse()

  r.PathPrefix("/static/").Handler(
    http.StripPrefix("/static/", http.FileServer(http.Dir(dir))),
  )

  routes.BackEndRoutes(r)
  routes.FrontEndRoutes(r)
}
