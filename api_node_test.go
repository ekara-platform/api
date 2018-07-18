package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNodes(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(getNodes)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestGetNodeDetails(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(getNodeDetails)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL + "/nodes/12")
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}
