package router

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/felunka/Mini-URL-Shortener-API/handler"
)

func SetupRouter() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)

  // Register API routes
  r.HandleFunc("/shorten", handler.ShortenURL).Methods(http.MethodPost, http.MethodOptions)
  log.Println("Route registered: POST /shorten")

  r.HandleFunc("/{shortURL}", handler.ResolveURL).Methods(http.MethodGet)
  log.Println("Route registered: GET /{shortURL}")

  return r
}
