package main

import (
	"github.com/gorilla/mux"
	"github.com/vavarodrigues/api/cmd/rest/handler"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")

	r.HandleFunc("/", handler.RedisHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
