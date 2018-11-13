package secret

type MockEmpySecret struct {
}

func (ms *MockEmpySecret) Store(key string, value []byte) error {
	return nil
}

func (ms *MockEmpySecret) StoreString(key string, value string) error {
	return nil
}

func (ms *MockEmpySecret) Get(key string) (bool, []byte, error) {
	return false, []byte(""), nil
}

func (ms *MockEmpySecret) Contains(key string) (bool, error) {
	return false, nil
}

func (ms *MockEmpySecret) Delete(key string) (bool, error) {
	return false, nil
}

func (ms *MockEmpySecret) Keys() ([]string, error) {
	r := make([]string, 0)
	return r, nil
}

func (ms *MockEmpySecret) Clean() error {
	return nil
}

type mockSecret struct {
	content map[string]string
}

func GetMockSecret() *mockSecret {
	return &mockSecret{
		content: make(map[string]string),
	}
}

func (ms *mockSecret) StoreSecret(key string, value []byte) error {
	ms.content[key] = string(value)
	return nil
}

func (ms *mockSecret) StoreSecretString(key string, value string) error {
	ms.content[key] = value
	return nil
}

func (ms *mockSecret) GetSecret(key string) (bool, []byte, error) {
	if val, ok := ms.content[key]; ok {
		return true, []byte(val), nil
	} else {
		return false, []byte(""), nil
	}
}

func (ms *mockSecret) ContainsSecret(key string) (bool, error) {
	if _, ok := ms.content[key]; ok {
		return true, nil
	}
	return false, nil
}

func (ms *mockSecret) DeleteSecret(key string) (bool, error) {
	delete(ms.content, key)
	return true, nil
}

func (ms *mockSecret) SecretKeys() ([]string, error) {
	r := make([]string, 0)
	for k, _ := range ms.content {
		r = append(r, k)
	}
	return r, nil
}

func (ms *mockSecret) CleanSecrets() error {
	keys, err := ms.SecretKeys()
	if err != nil {
		return err
	}

	for _, v := range keys {
		_, err := ms.DeleteSecret(v)
		if err != nil {
			return err
		}
	}
	return nil
}
