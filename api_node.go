package api

import (
	_ "encoding/json"
	"net/http"
)

// getNodes returns the nodes available into the environment
func getNodes(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
}

// getNodeDetails returns the details of the node corresponding
// to the id received as parameter
func getNodeDetails(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
}
