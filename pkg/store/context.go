package store

import (
	"context"
)

const key = "store"

func FromContext(ctx context.Context) *Store {
	return ctx.Value(key).(*Store)
}

func ToContext(ctx context.Context, s *Store) context.Context {
	return context.WithValue(ctx, key, s)
}
