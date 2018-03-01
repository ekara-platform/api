package main

import (
	"net/http"
	"testing"
)

func TestGetTasks(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tasks/", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestGetTaskDetails(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tasks/12", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestLaunchTask(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/tasks/12", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}
