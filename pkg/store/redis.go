package store

import (
	"fmt"
)

const queue = "queue"

func (s *Store) GetKey() (string, error) {

	pong, err := s.redis.Ping().Result()
	if err != nil {
		return "", err
	}
	fmt.Println(pong, err)

	return "teste", nil
}

func (s *Store) Push(p []byte) error {
	fmt.Println("Payload received: ")

	err := s.redis.RPush(queue, p).Err()
	if err != nil {
		return err
	}

	return nil
}
