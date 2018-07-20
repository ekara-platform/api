package consul

import (
	cApi "github.com/hashicorp/consul/api"
)

type ConsulStorage struct {
	client *cApi.Client
}

func Storage() (ConsulStorage, error) {
	r := &ConsulStorage{}
	client, err := cApi.NewClient(cApi.DefaultConfig())
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
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		return false, []byte(""), err
	}
	return true, pair.Value, nil
}

func (r ConsulStorage) Contains(key string) (bool, error) {
	kv := r.client.KV()
	pair, _, err := kv.Get("foo", nil)
	if err != nil {
		return false, err
	}
	res := false
	if pair != nil {
		res = true
	}

	return res, nil
}
