package main

import (
	_ "encoding/json"
	"net/http"
)

// getTasks returns the tasks available into the environment
func getTasks(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
}

// getTaskDetails returns the details of the task corresponding
// to the id received as parameter
func getTaskDetails(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
}

// runTask runs the task corresponding
// to the id received as parameter
func runTask(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
}
