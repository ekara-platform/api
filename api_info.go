package api

import (
	"encoding/json"
	"net/http"
)

type ApiInfo struct {
	Version             string
	Url                 string
	EnvironmentLocation string
	EnvironmentTime     string
	EnvironmentParam    string
	EnvironmentJSON     string
	EnvironmentYAML     string
	Err                 string
}

// getInfo returns the application detailed informations
func getInfo(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()
	w.Header().Set("Content-Type", MimeTypeJSON)

	storage := getStorage()

	info := ApiInfo{
		Version: application.Version,
		Url:     r.Host,
	}

	_, val, err := storage.Get(KEY_STORE_ENVIRONMENT_LOCATION)
	if err != nil {
		info.EnvironmentLocation = err.Error()
	}
	info.EnvironmentLocation = string(val)

	_, val, err = storage.Get(KEY_STORE_ENVIRONMENT_UPLOAD_TIME)
	if err != nil {
		info.EnvironmentTime = err.Error()
	}
	info.EnvironmentTime = string(val)

	_, val, err = storage.Get(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
	if err != nil {
		info.EnvironmentParam = err.Error()
	}
	info.EnvironmentParam = string(val)

	_, val, err = storage.Get(KEY_STORE_ENVIRONMENT_JSON_CONTENT)
	if err != nil {
		info.EnvironmentJSON = err.Error()
	}
	info.EnvironmentJSON = string(val)

	_, val, err = storage.Get(KEY_STORE_ENVIRONMENT_YAML_CONTENT)
	if err != nil {
		info.EnvironmentYAML = err.Error()
	}
	info.EnvironmentYAML = string(val)

	infoJSON, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	TResult.Print(string(infoJSON))
	w.Write(infoJSON)
}
