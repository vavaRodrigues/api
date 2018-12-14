package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

// NewClient newClient
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

// Ping ping
func Ping() error {
	client := NewClient()

	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}
