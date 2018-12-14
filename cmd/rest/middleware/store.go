package middleware

import (
	"net/http"
	"os"

	"github.com/vavarodrigues/api/pkg/store"
)

type StoreMiddleware struct {
	Store *store.Store
}

func NewStoreMiddleware() *StoreMiddleware {
	store := store.NewStore(
		store.StoreSettings{
			Addr:     os.Getenv("ADDR"),
			Password: os.Getenv("PASSWORD"),
			DB:       0,
		},
	)

	return &StoreMiddleware{Store: store}
}

func (m *StoreMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(store.ToContext(r.Context(), m.Store)))
	})
}
