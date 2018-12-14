package handler

import (
	"fmt"
	"github.com/vavarodrigues/api/pkg/store"
	"io/ioutil"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func PushHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := store.FromContext(r.Context()).Push(body); err == nil {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	s := store.FromContext(r.Context())

	pong, err := s.GetKey()
	if err != nil {
		fmt.Print("deu ruim")
	}
	fmt.Println(pong, err)

	fmt.Fprintf(w, "~> Redis %s", pong)

}
