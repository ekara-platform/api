package api

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestGetNodes(t *testing.T) {
	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/nodes/", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotImplemented, respRecorder.Code)
	checkEmptyRecoder(t, respRecorder)
}

func TestGetNodeDetails(t *testing.T) {
	application = App{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/nodes/12", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotImplemented, respRecorder.Code)
	checkEmptyRecoder(t, respRecorder)
}
