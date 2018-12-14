package main

import (
	"github.com/gorilla/mux"
	"github.com/vavarodrigues/api/cmd/rest/handler"
	"github.com/vavarodrigues/api/cmd/rest/middleware"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")
	r.HandleFunc("/", handler.GetKeyHandler).Methods("POST")
	r.HandleFunc("/push", handler.PushHandler).Methods("POST")

	storeMiddleware := middleware.NewStoreMiddleware()
	r.Use(storeMiddleware.Middleware)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
