package storage

import (
	"strings"

	"github.com/lagoon-platform/api/consul"
)

const (
	LAGOON_PREFIX string = "lagoon_"

	KEY_STORE_ENV_LOCATION    string = LAGOON_PREFIX + "environment_location"
	KEY_STORE_ENV_YAML        string = LAGOON_PREFIX + "environment_yaml_content"
	KEY_STORE_ENV_JSON        string = LAGOON_PREFIX + "environment_json_content"
	KEY_STORE_ENV_CREATED_AT  string = LAGOON_PREFIX + "environment_created_at"
	KEY_STORE_ENV_UPDATED_AT  string = LAGOON_PREFIX + "environment_updated_at"
	KEY_STORE_ENV_PARAM       string = LAGOON_PREFIX + "environment_param_content"
	KEY_STORE_ENV_SESSION     string = LAGOON_PREFIX + "environment_session_content"
	KEY_STORE_ENV_SSH_PRIVATE string = LAGOON_PREFIX + "environment_ssh_private"
	KEY_STORE_ENV_SSH_PUBLIC  string = LAGOON_PREFIX + "environment_ssh_public"
)

func RemoveLagoonPrefix(s string) string {
	if i := strings.Index(s, LAGOON_PREFIX); i == 0 {
		t := strings.Split(s, LAGOON_PREFIX)
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
}

func GetConsulStorage() Storage {
	s, err := consul.Storage()
	if err != nil {
		panic(err)
	}
	return s
}
