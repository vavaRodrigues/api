package handler

import (
	"fmt"
	"github.com/vavarodrigues/api/pkg/redis"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there")
}

func RedisHandler(w http.ResponseWriter, r *http.Request) {
	redis.Ping()

	fmt.Fprintf(w, "~> Redis %s", redis.Ping())
}
