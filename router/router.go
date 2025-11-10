package router

import (
  "log"

  "github.com/gorilla/mux"
  "github.com/felunka/Mini-URL-Shortener-API/handler"
)

func SetupRouter() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)

  // Register API routes
  r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
  log.Println("Route registered: POST /shorten")

  r.HandleFunc("/{shortURL}", handler.ResolveURL).Methods("GET")
  log.Println("Route registered: GET /{shortURL}")

  return r
}
