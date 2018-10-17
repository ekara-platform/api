package storage

import (
	"strings"

	"github.com/ekara-platform/api/consul"
)

const (
	EKARA_PREFIX string = "ekara_"

	KEY_STORE_ENV_LOCATION    string = EKARA_PREFIX + "environment_location"
	KEY_STORE_ENV_YAML        string = EKARA_PREFIX + "environment_yaml_content"
	KEY_STORE_ENV_JSON        string = EKARA_PREFIX + "environment_json_content"
	KEY_STORE_ENV_CREATED_AT  string = EKARA_PREFIX + "environment_created_at"
	KEY_STORE_ENV_UPDATED_AT  string = EKARA_PREFIX + "environment_updated_at"
	KEY_STORE_ENV_PARAM       string = EKARA_PREFIX + "environment_param_content"
	KEY_STORE_ENV_SESSION     string = EKARA_PREFIX + "environment_session_content"
	KEY_STORE_ENV_SSH_PRIVATE string = EKARA_PREFIX + "environment_ssh_private"
	KEY_STORE_ENV_SSH_PUBLIC  string = EKARA_PREFIX + "environment_ssh_public"
)

func RemoveEkaraPrefix(s string) string {
	if i := strings.Index(s, EKARA_PREFIX); i == 0 {
		t := strings.Split(s, EKARA_PREFIX)
		return t[1]
	}
	return s
}

type Storage interface {
	Store(key string, value []byte) error
	StoreString(key string, value string) error
	Get(key string) (bool, []byte, error)
	Contains(key string) (bool, error)
	Delete(key string) (bool, error)
	Keys() ([]string, error)
	Clean() error
}

func GetConsulStorage() Storage {
	s, err := consul.Storage()
	if err != nil {
		panic(err)
	}
	return s
}
