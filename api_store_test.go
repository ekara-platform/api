package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/ekara-platform/api/storage"
	"github.com/stretchr/testify/assert"
)

//************************************************
// EMPTY STORAGE WITHOUT ANY KEYS - START
//************************************************
func TestGetNoContent(t *testing.T) {

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/storage/"+"dummy_id", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

func TestDeleteNoContent(t *testing.T) {

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodDelete, "/storage/"+"dummy_id", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

func TestGetKeysNoContent(t *testing.T) {
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/storage/", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusOK, respRecorder.Code)

	keys := make([]string, 0)
	err := json.Unmarshal(respRecorder.Body.Bytes(), &keys)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, len(keys), 0)

}

//************************************************
// EMPTY STORAGE WITHOUT ANY KEYS - END
//************************************************

func TestSaveValue(t *testing.T) {

	strKey := "test_key"
	strValue := "test_value"

	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	body := StorePostRequest{
		Key:   strKey,
		Value: strValue,
	}

	jsonStr, err := json.Marshal(body)
	assert.Nil(t, err)

	req, _ := http.NewRequest(http.MethodPost, "/storage/", bytes.NewBuffer(jsonStr))
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusCreated, respRecorder.Code)

	b, err := usedStorage.Contains(strKey)
	assert.Nil(t, err)
	assert.True(t, b)
	b, val, err := usedStorage.Get(strKey)
	assert.Nil(t, err)
	assert.True(t, b)
	assert.Equal(t, string(val), strValue)

}

func TestGetKeys(t *testing.T) {
	usedStorage = storage.GetMockStorage()
	defer usedStorage.Clean()

	strKey1 := "test_key1"
	strValue1 := "test_value1"
	strKey2 := "test_key2"
	strValue2 := "test_value2"
	strKey3 := "test_key3"
	strValue3 := "test_value3"

	usedStorage.StoreString(strKey1, strValue1)
	usedStorage.StoreString(strKey2, strValue2)
	usedStorage.StoreString(strKey3, strValue3)

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/storage/", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusOK, respRecorder.Code)

	keys := make([]string, 0)
	err := json.Unmarshal(respRecorder.Body.Bytes(), &keys)
	if err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)
	assert.Contains(t, keys, strKey1, strKey2, strKey3)
}

func TestGetValue(t *testing.T) {

	strKey := "test_key"
	strValue := "test_value"

	usedStorage = storage.GetMockStorage()
	usedStorage.StoreString(strKey, strValue)
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/storage/"+strKey, nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusOK, respRecorder.Code)

	sPr := StorePostRequest{}
	err := json.Unmarshal(respRecorder.Body.Bytes(), &sPr)
	assert.Nil(t, err)
	assert.Equal(t, sPr.Key, strKey)
	assert.Equal(t, sPr.Value, strValue)

}

func TestDeleteValue(t *testing.T) {

	strKey := "test_key"
	strValue := "test_value"

	usedStorage = storage.GetMockStorage()
	usedStorage.StoreString(strKey, strValue)
	defer usedStorage.Clean()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodDelete, "/storage/"+strKey, nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusOK, respRecorder.Code)

	b, err := usedStorage.Contains(strKey)

	assert.Nil(t, err)
	assert.False(t, b)

}
