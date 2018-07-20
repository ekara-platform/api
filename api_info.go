package api

import (
	"encoding/json"
	"net/http"
)

type ApiInfo struct {
	Version string
	Url     string
	Key1    string
	Err     string
}

// getInfo returns the application detailed informations
func getInfo(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", "application/json")

	storage := getStorage()
	storage.StoreString("key1", "value")

	info := ApiInfo{
		application.Version,
		r.Host,
		"No value",
		"No error",
	}

	_, err := storage.Contains("key1")
	if err != nil {
		info.Err = err.Error()
	}

	infoJSON, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	TResult.Print(string(infoJSON))
	w.Write(infoJSON)
}
