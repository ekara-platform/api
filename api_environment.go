package main

import (
	"encoding/json"
	"net/http"
)

type EnvironmentLoadRequest struct {
	Location string `json:"location"`
}

// getEnvironment returns the details enviroment currently
// manage by lagoon
func getEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
}

// loadEnvironment loads the enviroment received as parameter
func loadEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var e EnvironmentLoadRequest
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO Parse the file content here

	TResult.Printf(ENVIRONMENT_CREATED, e.Location)
	w.WriteHeader(http.StatusCreated)
}

// deleteEnvironment deletes the enviroment currently
// manage by lagoon
func deleteEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status

	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
}
