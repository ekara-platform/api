package api

import (
	"github.com/lagoon-platform/api/consul"
)

const (
	KEY_STORE_ENVIRONMENT_LOCATION      string = "environment_location"
	KEY_STORE_ENVIRONMENT_UPLOAD_TIME   string = "environment_upload_time"
	KEY_STORE_ENVIRONMENT_PARAM_CONTENT string = "environment_param_content"
)

type Storage interface {
	Store(key string, value []byte) error
	StoreString(key string, value string) error
	Get(key string) (bool, []byte, error)
	Contains(key string) (bool, error)
	Delete(key string) (bool, error)
}

func getStorage() Storage {
	s, err := consul.Storage()
	if err != nil {
		panic(err)
	}
	return s
}
