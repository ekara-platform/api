package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/ekara-platform/api/secret"
	"github.com/stretchr/testify/assert"
)

//************************************************
// EMPTY SECRET WITHOUT ANY KEYS - START
//************************************************
func TestGetSecretNoContent(t *testing.T) {
	usedSecret = secret.GetMockSecret()
	defer usedSecret.CleanSecrets()

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/secret/"+"dummy_id", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

func TestDeleteSecretNoContent(t *testing.T) {
	usedSecret = secret.GetMockSecret()
	defer usedSecret.CleanSecrets()

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodDelete, "/secret/"+"dummy_id", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotFound, respRecorder.Code)

}

func TestGetSecretKeysNoContent(t *testing.T) {
	usedSecret = secret.GetMockSecret()
	defer usedSecret.CleanSecrets()

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/secret/", nil)
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
// EMPTY SECRET WITHOUT ANY KEYS - END
//************************************************

func TestSaveSecretValue(t *testing.T) {

	strKey := "test_key"
	strValue := "test_value"

	usedSecret = secret.GetMockSecret()
	defer usedSecret.CleanSecrets()

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

	req, _ := http.NewRequest(http.MethodPost, "/secret/", bytes.NewBuffer(jsonStr))
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusCreated, respRecorder.Code)

	b, err := usedSecret.ContainsSecret(strKey)
	assert.Nil(t, err)
	assert.True(t, b)
	b, val, err := usedSecret.GetSecret(strKey)
	assert.Nil(t, err)
	assert.True(t, b)
	assert.Equal(t, string(val), strValue)

}

func TestGetSecretKeys(t *testing.T) {
	usedSecret = secret.GetMockSecret()
	defer usedSecret.CleanSecrets()

	strKey1 := "test_key1"
	strValue1 := "test_value1"
	strKey2 := "test_key2"
	strValue2 := "test_value2"
	strKey3 := "test_key3"
	strValue3 := "test_value3"

	usedSecret.StoreSecretString(strKey1, strValue1)
	usedSecret.StoreSecretString(strKey2, strValue2)
	usedSecret.StoreSecretString(strKey3, strValue3)

	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/secret/", nil)
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

func TestGetSecretValue(t *testing.T) {

	strKey := "test_key"
	strValue := "test_value"

	usedSecret = secret.GetMockSecret()
	usedSecret.StoreSecretString(strKey, strValue)
	defer usedSecret.CleanSecrets()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodGet, "/secret/"+strKey, nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusOK, respRecorder.Code)

	sPr := StorePostRequest{}
	err := json.Unmarshal(respRecorder.Body.Bytes(), &sPr)
	assert.Nil(t, err)
	assert.Equal(t, sPr.Key, strKey)
	assert.Equal(t, sPr.Value, strValue)

}

func TestDeleteSecretValue(t *testing.T) {

	strKey := "test_key"
	strValue := "test_value"

	usedSecret = secret.GetMockSecret()
	usedSecret.StoreSecretString(strKey, strValue)
	defer usedSecret.CleanSecrets()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	application = App{}
	application.initialize()

	req, _ := http.NewRequest(http.MethodDelete, "/secret/"+strKey, nil)
	respRecorder := executeRequest(req)
	checkCode(t, http.StatusOK, respRecorder.Code)

	b, err := usedSecret.ContainsSecret(strKey)

	assert.Nil(t, err)
	assert.False(t, b)

}
