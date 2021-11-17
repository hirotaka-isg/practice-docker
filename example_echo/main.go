package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Println("received request")
    fmt.Fprintf(w, "Hello Docker!!")
  })

  log.Println("start server")
  server := &http.server{Addr: ":8000"}
  if err := server.ListernAndServer(); err != nil {
    log.Println(err)
  }
}
