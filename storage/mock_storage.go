package storage

type MockEmpyStorage struct {
}

func (ms *MockEmpyStorage) Store(key string, value []byte) error {
	return nil
}

func (ms *MockEmpyStorage) StoreString(key string, value string) error {
	return nil
}

func (ms *MockEmpyStorage) Get(key string) (bool, []byte, error) {
	return false, []byte(""), nil
}

func (ms *MockEmpyStorage) Contains(key string) (bool, error) {
	return false, nil
}

func (ms *MockEmpyStorage) Delete(key string) (bool, error) {
	return false, nil
}

func (ms *MockEmpyStorage) Keys() ([]string, error) {
	r := make([]string, 0)
	return r, nil
}

type mockStorage struct {
	content map[string]string
}

func GetMockStorage() *mockStorage {
	return &mockStorage{
		content: make(map[string]string),
	}
}

func (ms *mockStorage) Store(key string, value []byte) error {
	ms.content[key] = string(value)
	return nil
}

func (ms *mockStorage) StoreString(key string, value string) error {
	ms.content[key] = value
	return nil
}

func (ms *mockStorage) Get(key string) (bool, []byte, error) {
	if val, ok := ms.content[key]; ok {
		return true, []byte(val), nil
	} else {
		return false, []byte(""), nil
	}
}

func (ms *mockStorage) Contains(key string) (bool, error) {
	if _, ok := ms.content[key]; ok {
		return true, nil
	}
	return false, nil
}

func (ms *mockStorage) Delete(key string) (bool, error) {
	delete(ms.content, key)
	return true, nil
}

func (ms *mockStorage) Keys() ([]string, error) {
	r := make([]string, 0)
	for k, _ := range ms.content {
		r = append(r, k)
	}
	return r, nil
}
