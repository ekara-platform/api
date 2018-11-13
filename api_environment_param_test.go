package api

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/ekara-platform/api/storage"
	"github.com/stretchr/testify/assert"
)

//************************************************
// EMPTY STORAGE WITHOUT ANY PARAMS - START
//************************************************
func TestGetNoParam(t *testing.T) {

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean(storage.STORAGE_PREFIX)

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/envparam/", nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

func TestDeleteNoParam(t *testing.T) {

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean(storage.STORAGE_PREFIX)

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodDelete, "/envparam/", nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

//************************************************
// EMPTY STORAGE WITHOUT ANY PARAMS - END
//************************************************

func TestSaveParam(t *testing.T) {

	strContent := "Dummy params content"

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean(storage.STORAGE_PREFIX)

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodPost, "/envparam/", bytes.NewBuffer([]byte(strContent)))
	respRecorder := executeRequest(req)
	resp := respRecorder.Result()

	checkResponseCode(t, http.StatusCreated, resp)

	b, err := usedStorage.Contains(storage.KEY_STORE_ENV_PARAM)
	assert.Nil(t, err)
	assert.True(t, b)
	b, val, err := usedStorage.Get(storage.KEY_STORE_ENV_PARAM)
	assert.Nil(t, err)
	assert.True(t, b)
	assert.Equal(t, string(val), strContent)

}

func TestGetParam(t *testing.T) {

	strContent := "Dummy params content"

	usedStorage = storage.GetMockStorage()
	usedStorage.StoreString(storage.KEY_STORE_ENV_PARAM, strContent)
	defer usedStorage.Clean(storage.STORAGE_PREFIX)

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/envparam/", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusOK, respRecorder.Code)

	assert.Equal(t, strContent, string(respRecorder.Body.Bytes()))
}

func TestDeleteParam(t *testing.T) {

	strContent := "Dummy params content"

	usedStorage = storage.GetMockStorage()
	usedStorage.StoreString(storage.KEY_STORE_ENV_PARAM, strContent)
	defer usedStorage.Clean(storage.STORAGE_PREFIX)

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest("DELETE", "/envparam/", nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusOK, respRecorder.Code)

	b, err := usedStorage.Contains(storage.KEY_STORE_ENV_PARAM)

	assert.Nil(t, err)
	assert.False(t, b)

}
