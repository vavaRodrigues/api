package handler

import (
	"fmt"
	"github.com/vavarodrigues/api/pkg/store"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
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
