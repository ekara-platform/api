package api

import (
	"github.com/lagoon-platform/api/consul"
)

type Storage interface {
	Store(key string, value []byte) error
	StoreString(key string, value string) error
	Get(key string) (bool, []byte, error)
	Contains(key string) (bool, error)
}

func getStorage() Storage {
	s, err := consul.Storage()
	if err != nil {
		panic(err)
	}
	return s
}
