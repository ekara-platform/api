package rest

import (
	"fmt"
	"net/http"
)

func FilterKeyFound(key, message string, w http.ResponseWriter) (callerShouldReturn bool) {
	s := api.storageEngine
	b, err := s.Contains(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		callerShouldReturn = true
		return
	}
	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be located into the storage", message)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		callerShouldReturn = true
		return
	}
	return
}

func FilterSecretKeyFound(key, message string, w http.ResponseWriter) (callerShouldReturn bool) {
	s := api.secretEngine
	b, err := s.ContainsSecret(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		callerShouldReturn = true
		return
	}
	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be located into the secrets", message)
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		callerShouldReturn = true
		return
	}
	return
}
