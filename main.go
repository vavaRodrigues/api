package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vavarodrigues/cmd/rest/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/", redisHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func redisHandler(w http.ResponseWriter, r *http.Request) {
	ping()

	fmt.Fprintf(w, "~> Redis %s", ping())
}

func ping() error {
	client := NewClient()

	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}
