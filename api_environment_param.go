package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getEnvironmentParam(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	b, val, err := getStorage().Get(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	TResult.Print(string(val))
	w.Header().Set("Content-Type", MimeTypeYAML)
	w.Write(val)
	w.WriteHeader(http.StatusOK)
}

func saveEnvironmentParam(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = getStorage().Store(KEY_STORE_ENVIRONMENT_PARAM_CONTENT, b)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO Refresh the environment if present

	TResult.Printf(VALUE_STORED, KEY_STORE_ENVIRONMENT_PARAM_CONTENT, string(b))
	w.WriteHeader(http.StatusCreated)
}

func deleteEnvironmentParam(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	s := getStorage()

	b, err := s.Contains(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err = getStorage().Delete(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if b {
		TResult.Printf(VALUE_DELETED, KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
	}
	w.WriteHeader(http.StatusOK)
}
