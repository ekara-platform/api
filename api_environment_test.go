package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEnvironment(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(getEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestDeleteEnvironment(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(deleteEnvironment)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}
