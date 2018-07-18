package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasks(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(getTasks)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestGetTaskDetails(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(getTaskDetails)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks/12")
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}

func TestLaunchTask(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(runTask)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL + "/tasks/12")
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusNotImplemented, resp)
	checkEmptyBody(t, resp)
}
