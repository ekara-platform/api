package rest

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestGetTasks(t *testing.T) {
	application = Api{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotImplemented, respRecorder.Code)
	checkEmptyRecoder(t, respRecorder)
}

func TestGetTaskDetails(t *testing.T) {
	application = Api{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/12", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotImplemented, respRecorder.Code)
	checkEmptyRecoder(t, respRecorder)
}

func TestLaunchTask(t *testing.T) {
	application = Api{}
	application.initialize()

	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)

	req, _ := http.NewRequest(http.MethodPut, "/tasks/12", nil)
	respRecorder := executeRequest(req)

	checkCode(t, http.StatusNotImplemented, respRecorder.Code)
	checkEmptyRecoder(t, respRecorder)

}
