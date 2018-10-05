package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/lagoon-platform/api/storage"
	"github.com/lagoon-platform/engine"
	"github.com/stretchr/testify/assert"
)

//************************************************
// EMPTY STORAGE WITHOUT ANY SESSION - START
//************************************************
func TestGetNoSession(t *testing.T) {

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/envsession/", nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

func TestDeleteNoSession(t *testing.T) {

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodDelete, "/envsession/", nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

//************************************************
// EMPTY STORAGE WITHOUT ANY SESSION - END
//************************************************

func TestSaveSession(t *testing.T) {

	es := engine.EngineSession{
		CreationSession: &engine.CreationSession{},
		File:            "file_path",
	}

	es.CreationSession.Client = "session_client"
	es.CreationSession.Uids = make(map[string]string)
	es.CreationSession.Uids["ID1"] = "ID1_content"
	es.CreationSession.Uids["ID2"] = "ID2_content"

	by, err := json.Marshal(es)
	assert.Nil(t, err)

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodPost, "/envsession/", bytes.NewBuffer(by))
	respRecorder := executeRequest(req)
	resp := respRecorder.Result()

	checkResponseCode(t, http.StatusCreated, resp)

	b, err := usedStorage.Contains(storage.KEY_STORE_ENV_SESSION)
	assert.Nil(t, err)
	assert.True(t, b)
	b, val, err := usedStorage.Get(storage.KEY_STORE_ENV_SESSION)
	assert.Nil(t, err)
	assert.True(t, b)
	assert.Equal(t, string(val), string(by))

}

func TestGetSession(t *testing.T) {

	strContent := "Dummy session content"

	usedStorage = storage.GetMockStorage()
	usedStorage.StoreString(storage.KEY_STORE_ENV_SESSION, strContent)
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/envsession/", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusOK, respRecorder.Code)

	assert.Equal(t, strContent, string(respRecorder.Body.Bytes()))
}

func TestDeleteSession(t *testing.T) {

	strContent := "Dummy session content"

	usedStorage = storage.GetMockStorage()
	usedStorage.StoreString(storage.KEY_STORE_ENV_SESSION, strContent)
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest("DELETE", "/envsession/", nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusOK, respRecorder.Code)

	b, err := usedStorage.Contains(storage.KEY_STORE_ENV_SESSION)

	assert.Nil(t, err)
	assert.False(t, b)

}
