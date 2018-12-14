package store

import (
	"fmt"
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
	fmt.Printf("%+v\n", settings)
	client := redis.NewClient(&redis.Options{
		Addr:     settings.Addr,
		Password: settings.Password, // no password set
		DB:       0,                 // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return &Store{redis: client}
}
