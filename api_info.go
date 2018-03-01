package main

import (
	"encoding/json"
	"net/http"
)

type ApiInfo struct {
	Version string
	Url     string
}

// getInfo returns the application detailed informations
func getInfo(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	info := ApiInfo{
		application.Version,
		r.Host,
	}

	infoJSON, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	TResult.Print(string(infoJSON))
	w.Write(infoJSON)
}
