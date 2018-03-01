package main

import (
	"net/http"
	"testing"
)

func TestGetNodes(t *testing.T) {
	req, _ := http.NewRequest("GET", "/nodes/", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestGetNodeDetails(t *testing.T) {
	req, _ := http.NewRequest("GET", "/nodes/12", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}
