package store

import (
	"github.com/go-redis/redis"
)

type Store struct {
	redis *redis.Client
}

type StoreSettings struct {
	Addr     string
	Password string
	DB       int
}

func NewStore(settings StoreSettings) *Store {
	client := redis.NewClient(&redis.Options{
		Addr:     settings.Addr,
		Password: settings.Password, // no password set
		DB:       settings.DB,       // use default DB
	})

	return &Store{redis: client}
}
