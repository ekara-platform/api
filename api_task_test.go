package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetTasks(t *testing.T) {
	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)
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
	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)
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
	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)
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
