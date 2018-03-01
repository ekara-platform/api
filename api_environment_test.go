package main

import (
	"net/http"
	"testing"
)

func TestGetEnvironment(t *testing.T) {
	req, _ := http.NewRequest("GET", "/environment/", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestDeleteEnvironment(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/environment/", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}
