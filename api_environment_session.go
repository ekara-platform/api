package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lagoon-platform/api/storage"
	"github.com/lagoon-platform/engine"
)

func getEnvironmentSession(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	if FilterKeyFound(storage.KEY_STORE_ENV_SESSION, "Session", w) {
		return
	}

	b, val, err := usedStorage.Get(storage.KEY_STORE_ENV_SESSION)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", storage.KEY_STORE_ENV_SESSION)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	TResult.Print(string(val))
	w.Header().Set("Content-Type", MimeTypeJSON)
	w.Write(val)
	w.WriteHeader(http.StatusOK)
}

func saveEnvironmentSession(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Check if the received content is a valid session
	session := &engine.EngineSession{}
	err = json.Unmarshal(b, session)
	if err != nil {
		err := fmt.Errorf(ERROR_CONTENT, "Reading the session content", err.Error())
		TLog.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(session.CreationSession.Client) == 0 {
		err := fmt.Errorf(ERROR_CONTENT, "Session client cannot be empty", err.Error())
		TLog.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(session.CreationSession.Uids) == 0 {
		err := fmt.Errorf(ERROR_CONTENT, "Session Uids cannot be empty", err.Error())
		TLog.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = usedStorage.Store(storage.KEY_STORE_ENV_SESSION, b)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO Refresh the environment if present

	TResult.Printf(VALUE_STORED, storage.KEY_STORE_ENV_SESSION, string(b))
	w.WriteHeader(http.StatusCreated)
}

func deleteEnvironmentSession(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	s := usedStorage
	if FilterKeyFound(storage.KEY_STORE_ENV_SESSION, "Session", w) {
		return
	}

	b, err := s.Contains(storage.KEY_STORE_ENV_SESSION)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", storage.KEY_STORE_ENV_SESSION)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err = usedStorage.Delete(storage.KEY_STORE_ENV_SESSION)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if b {
		TResult.Printf(VALUE_DELETED, storage.KEY_STORE_ENV_SESSION)
	}
	w.WriteHeader(http.StatusOK)
}
