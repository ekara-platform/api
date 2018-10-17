package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ekara-platform/api/storage"
)

func getEnvironmentParam(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	b, val, err := usedStorage.Get(storage.KEY_STORE_ENV_PARAM)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", storage.KEY_STORE_ENV_PARAM)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
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

	err = usedStorage.Store(storage.KEY_STORE_ENV_PARAM, b)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO Refresh the environment if present

	TResult.Printf(VALUE_STORED, storage.KEY_STORE_ENV_PARAM, string(b))
	w.WriteHeader(http.StatusCreated)
}

func deleteEnvironmentParam(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	s := usedStorage

	b, err := s.Contains(storage.KEY_STORE_ENV_PARAM)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", storage.KEY_STORE_ENV_PARAM)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	b, err = usedStorage.Delete(storage.KEY_STORE_ENV_PARAM)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if b {
		TResult.Printf(VALUE_DELETED, storage.KEY_STORE_ENV_PARAM)
	}
	w.WriteHeader(http.StatusOK)
}
