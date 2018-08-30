package consul

import (
	cApi "github.com/hashicorp/consul/api"
)

type ConsulStorage struct {
	client *cApi.Client
}

func Storage() (ConsulStorage, error) {
	r := &ConsulStorage{}
	config := cApi.DefaultConfig()
	config.Address = "consul:8500"

	client, err := cApi.NewClient(config)
	if err != nil {
		return *r, err
	}
	r.client = client
	return *r, nil
}

func (r ConsulStorage) Store(key string, value []byte) error {
	kv := r.client.KV()

	p := &cApi.KVPair{Key: key, Value: value}
	_, err := kv.Put(p, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r ConsulStorage) StoreString(key string, value string) error {
	return r.Store(key, []byte(value))
}

func (r ConsulStorage) Get(key string) (bool, []byte, error) {
	kv := r.client.KV()
	pair, _, err := kv.Get(key, nil)

	if err != nil {
		return false, []byte(""), err
	}
	if pair == nil {
		return false, []byte(""), nil
	}
	return true, pair.Value, nil
}

func (r ConsulStorage) Contains(key string) (bool, error) {
	kv := r.client.KV()
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return false, err
	}
	if pair != nil {
		return true, nil
	}
	return false, nil
}

func (r ConsulStorage) Delete(key string) (bool, error) {
	kv := r.client.KV()
	_, err := kv.Delete(key, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
