package store

import (
	"fmt"
)

func (s *Store) GetKey() (string, error) {

	pong, err := s.redis.Ping().Result()
	if err != nil {
		return "", err
	}
	fmt.Println(pong, err)

	return "teste", nil
}
