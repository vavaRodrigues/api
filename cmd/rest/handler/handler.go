package handler

import (
	"fmt"
	"github.com/vavarodrigues/api/pkg/redis"
	"github.com/vavarodrigues/api/pkg/store"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func RedisHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(redis.Ping())

	fmt.Fprintf(w, "~> Redis %s", redis.Ping())
}

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	s := store.FromContext(r.Context())

	fmt.Fprintf(w, "~> Redis %s", s.GetKey)

}
