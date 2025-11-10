package main

import (
  "log"
  "net/http"

  "github.com/felunka/Mini-URL-Shortener-API/router"
)

func main() {
  r := router.SetupRouter()

  port := "8080"
  err := http.ListenAndServe(":"+port, r)
  if err != nil {
    log.Fatalf("Server failed to start: %v", err)
  } else {
    log.Println("Started!")
  }
}
