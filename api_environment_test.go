package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/environment/", nil)
	respRecorder := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, respRecorder.Result())
}

func TestPutNoEnvironment(t *testing.T) {
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodPut, "/environment/", bytes.NewBuffer([]byte("Dummy content")))
	respRecorder := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, respRecorder.Result())
}

func TestDeleteNoEnvironment(t *testing.T) {

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodDelete, "/environment/", nil)
	respRecorder := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, respRecorder.Result())
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

	application = App{}
	application.initialize()

	body := EnvironmentLoadRequest{
		Location: str,
	}
	jsonStr, err := json.Marshal(body)
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/environment/", bytes.NewBuffer(jsonStr))
	respRecorder := executeRequest(req)

	checkResponseCode(t, http.StatusConflict, respRecorder.Result())

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

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/environment/", nil)
	req.Header.Set("Content-type", MimeTypeYAML)
	respRecorder := executeRequest(req)
	resp := respRecorder.Result()

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

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/environment/", nil)
	req.Header.Set("Content-type", MimeTypeJSON)
	respRecorder := executeRequest(req)
	resp := respRecorder.Result()

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

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/environment/", nil)
	respRecorder := executeRequest(req)
	resp := respRecorder.Result()

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

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodDelete, "/environment/", nil)
	respRecorder := executeRequest(req)
	resp := respRecorder.Result()

	checkResponseCode(t, http.StatusOK, resp)
	keys, _ := usedStorage.Keys()
	assert.Equal(t, len(keys), 0)

}
