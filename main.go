package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/key/{key}", kvGetHandler).Methods("GET")
	r.HandleFunc("/v1/key/{key}", kvPutHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
