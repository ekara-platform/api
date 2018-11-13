package secret

import (
	"strings"

	"github.com/ekara-platform/api/consul"
)

const (
	SECRET_PREFIX string = "storage_"
)

type (
	Secret interface {
		StoreSecret(key string, value []byte) error
		StoreSecretString(key string, value string) error
		GetSecret(key string) (bool, []byte, error)
		ContainsSecret(key string) (bool, error)
		DeleteSecret(key string) (bool, error)
		SecretKeys() ([]string, error)
		CleanSecrets() error
	}

	ConsulBridge struct {
		secretSystem consul.ConsulStorage
	}
)

func (r ConsulBridge) StoreSecret(key string, value []byte) error {
	if !strings.HasPrefix("key", SECRET_PREFIX) {
		key = SECRET_PREFIX + key
	}
	return r.secretSystem.Store(key, value)
}

func (r ConsulBridge) StoreSecretString(key string, value string) error {
	if !strings.HasPrefix("key", SECRET_PREFIX) {
		key = SECRET_PREFIX + key
	}
	return r.secretSystem.StoreString(key, value)
}

func (r ConsulBridge) GetSecret(key string) (bool, []byte, error) {
	if !strings.HasPrefix("key", SECRET_PREFIX) {
		key = SECRET_PREFIX + key
	}
	return r.secretSystem.Get(key)
}

func (r ConsulBridge) ContainsSecret(key string) (bool, error) {
	if !strings.HasPrefix("key", SECRET_PREFIX) {
		key = SECRET_PREFIX + key
	}
	return r.secretSystem.Contains(key)
}

func (r ConsulBridge) DeleteSecret(key string) (bool, error) {
	if !strings.HasPrefix("key", SECRET_PREFIX) {
		key = SECRET_PREFIX + key
	}
	return r.secretSystem.Delete(key)
}

func (r ConsulBridge) SecretKeys() ([]string, error) {
	result := make([]string, 0)
	keys, err := r.secretSystem.Keys()
	if err != nil {
		return result, err
	}
	for _, v := range keys {
		if strings.HasPrefix(v, SECRET_PREFIX) {
			result = append(result, v[len(SECRET_PREFIX):])
		}
	}
	return result, nil
}

func (r ConsulBridge) CleanSecrets() error {
	return r.secretSystem.Clean(SECRET_PREFIX)
}

func GetSecret() Secret {
	// TODO Later remove consul and use something else once the technology has been identified
	s, err := consul.Storage()
	if err != nil {
		panic(err)
	}

	return ConsulBridge{secretSystem: s}
}
