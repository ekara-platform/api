package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"net/http/httptest"

	"testing"
)

// checkResponseCode checks if the request response code
// equals the expected one
func checkResponseCode(t *testing.T, expected int, resp *http.Response) {
	if expected != resp.StatusCode {
		t.Errorf("Expected response code %d. Got %d\n", expected, resp.StatusCode)
	}
}

// checkCode checks if the actual code
// equals the expected one
func checkCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// getNotEmptyBody returns the non empty body of the provided response
// If the body is empty the test will result in a Errorf,
func getNotEmptyBody(t *testing.T, resp *http.Response) string {
	var body []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) == "" {
		t.Errorf("Expected an body. Got nothing")
	}
	return string(body)
}

// checkEmptyBody checks if body of the provided response is empty
// If the body is not empty the test will result in a Errorf,
func checkEmptyBody(t *testing.T, resp *http.Response) {
	var body []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(body) > 0 {
		t.Errorf("Expected nothing. Got something")
	}
}

// checkEmptyRecoder checks if body of the provided recorder is empty
// If the body is not empty the test will result in a Errorf,
func checkEmptyRecoder(t *testing.T, rec *httptest.ResponseRecorder) {
	if len(rec.Body.Bytes()) > 0 {
		t.Errorf("Expected nothing. Got something")
	}
}

func checkBody(t *testing.T, resp *http.Response, expected []byte) {
	var body []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(body) == 0 {
		t.Errorf("Expected something. Got nothing")
	} else {
		eStr := strings.Trim(string(expected), " ")
		bStr := strings.Trim(string(body), " ")
		bStr = strings.Trim(string(bStr), "\t")

		if eStr != bStr {
			t.Errorf("Expected : \"%s\". Got : \"%s\"", expected, body)
		}
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

// TODO DOCUMENTATION
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	application.Router.ServeHTTP(rr, req)

	return rr
}
