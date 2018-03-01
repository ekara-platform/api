package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}

	initLog(false, false)
	a.Initialize()

	code := m.Run()

	os.Exit(code)
}

// checkResponseCode checks if the request response code
// equals the expected one
func checkResponseCode(t *testing.T, expected int, resp *httptest.ResponseRecorder) {
	if expected != resp.Code {
		t.Errorf("Expected response code %d. Got %d\n", expected, resp.Code)
	}
}

// executeRequest executes the given request against
// the application router
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	r := httptest.NewRecorder()
	a.Router.ServeHTTP(r, req)
	return r
}

// getNotEmptyBody returns the non empty body of the provided response
// If the body is empty the test will result in a Errorf,
func getNotEmptyBody(t *testing.T, resp *httptest.ResponseRecorder) string {
	var body string
	if body = resp.Body.String(); body == "" {
		t.Errorf("Expected an body. Got nothing")
	}
	return body
}

// checkEmptyBody checks if body of the provided response is empty
// If the body is not empty the test will result in a Errorf,
func checkEmptyBody(t *testing.T, resp *httptest.ResponseRecorder) {
	if body := resp.Body.String(); body != "" {
		t.Errorf("Expected an body. Got nothing")
	}
}

// checkJsonRoundTpip will Unmarshal and then Marshal again the json string
// using the given interface.
// We use this marshalling roud trip to validate that the received json is well
// suited to the given interface.
// If the round trip doesn't suffer data loss then it means that the interface/structure
// use to generate the received JSON is the same that the tested interface
//
// The check is done this way because the json.Unmarshal method
// won't throw an error in the input string cannot be unmarshal into the interface.
// If the jsonString cannot be Unmarshaled or Marshaled the test will result in a Errorf
func checkJsonRoundTpip(t *testing.T, jsonString string, across interface{}) (bool, string) {
	a := &across
	if err := json.Unmarshal([]byte(jsonString), a); err != nil {
		t.Errorf(err.Error())
	}

	var output []byte
	var err error
	if output, err = json.Marshal(a); err != nil {
		t.Errorf(err.Error())
	}
	// if the output is not equals to jsonString
	// then data loss
	return jsonString == string(output), string(output)
}
