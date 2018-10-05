package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/lagoon-platform/api/storage"
	"github.com/stretchr/testify/assert"
)

//************************************************
// EMPTY STORAGE WITHOUT ANY ENVIRONMENT - START
//************************************************
func TestGetNoEnvironment(t *testing.T) {
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	handler := http.HandlerFunc(getEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotFound, resp)
}

func TestPutNoEnvironment(t *testing.T) {
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	handler := http.HandlerFunc(updateEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotFound, resp)
}

func TestDeleteNoEnvironment(t *testing.T) {

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	handler := http.HandlerFunc(deleteEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	req := httptest.NewRequest(http.MethodDelete, server.URL, nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	checkResponseCode(t, http.StatusNotFound, resp)
}

//************************************************
// EMPTY STORAGE WITHOUT ANY ENVIRONMENT - END
//************************************************

//************************************************
// UNIQUE ENVIRONMENT INTO THE STORAGE- START
//************************************************

func TestPostSecondEnvironment(t *testing.T) {

	str := "dummy_environment_location"
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()
	usedStorage.StoreString(storage.KEY_STORE_ENV_LOCATION, str)

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	handler := http.HandlerFunc(loadEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	body := EnvironmentLoadRequest{
		Location: str,
	}
	jsonStr, err := json.Marshal(body)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, server.URL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	checkResponseCode(t, http.StatusConflict, resp)
}

//************************************************
// UNIQUE ENVIRONMENT INTO THE STORAGE- END
//************************************************

func TestGetEnvironmentYaml(t *testing.T) {

	str := "dummy_environment"
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()
	usedStorage.StoreString(storage.KEY_STORE_ENV_LOCATION, str+"_location")
	usedStorage.StoreString(storage.KEY_STORE_ENV_YAML, str+"_YAML")
	usedStorage.StoreString(storage.KEY_STORE_ENV_JSON, str+"_JSON")

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	handler := http.HandlerFunc(getEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	req := httptest.NewRequest(http.MethodGet, server.URL, nil)
	req.Header.Set("Content-type", MimeTypeYAML)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	checkResponseCode(t, http.StatusOK, resp)
	checkBody(t, resp, []byte(str+"_YAML"))
}

func TestGetEnvironmentJson(t *testing.T) {

	str := "dummy_environment"
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()
	usedStorage.StoreString(storage.KEY_STORE_ENV_LOCATION, str+"_location")
	usedStorage.StoreString(storage.KEY_STORE_ENV_YAML, str+"_YAML")
	usedStorage.StoreString(storage.KEY_STORE_ENV_JSON, str+"_JSON")

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	handler := http.HandlerFunc(getEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	req := httptest.NewRequest(http.MethodGet, server.URL, nil)
	req.Header.Set("Content-type", MimeTypeJSON)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	checkResponseCode(t, http.StatusOK, resp)
	checkBody(t, resp, []byte(str+"_JSON"))
}

func TestGetEnvironmentNoContentType(t *testing.T) {

	str := "dummy_environment"
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()
	usedStorage.StoreString(storage.KEY_STORE_ENV_LOCATION, str+"_location")
	usedStorage.StoreString(storage.KEY_STORE_ENV_YAML, str+"_YAML")
	usedStorage.StoreString(storage.KEY_STORE_ENV_JSON, str+"_JSON")

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	handler := http.HandlerFunc(getEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	req := httptest.NewRequest(http.MethodGet, server.URL, nil)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	checkResponseCode(t, http.StatusOK, resp)
	checkBody(t, resp, []byte(str+"_JSON"))
}

func TestDeleteEnvironment(t *testing.T) {

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()
	usedStorage.StoreString(storage.KEY_STORE_ENV_JSON, "KEY_STORE_ENV_JSON_CONTENT")
	usedStorage.StoreString(storage.KEY_STORE_ENV_LOCATION, "KEY_STORE_ENV_LOCATION_CONTENT")
	usedStorage.StoreString(storage.KEY_STORE_ENV_CREATED_AT, "KEY_STORE_ENV_CREATED_AT")
	usedStorage.StoreString(storage.KEY_STORE_ENV_UPDATED_AT, "KEY_STORE_ENV_UPDATED_AT")

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)
	handler := http.HandlerFunc(deleteEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	req := httptest.NewRequest(http.MethodDelete, server.URL, nil)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	checkResponseCode(t, http.StatusOK, resp)
	keys, _ := usedStorage.Keys()
	assert.Equal(t, len(keys), 0)

}
