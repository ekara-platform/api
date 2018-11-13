package api

import (
	"encoding/json"
	_ "fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// getSecret returns the secret content corresponding to the key received
// as parameter
func getSecret(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	vars := mux.Vars(r)
	id := vars["id"]
	if FilterSecretKeyFound(id, id, w) {
		return
	}

	_, val, err := usedSecret.GetSecret(id)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := StorePostRequest{
		id,
		string(val),
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	TResult.Print(string(resultJSON))
	w.Header().Set("Content-Type", MimeTypeJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

// getSecretKeys returns the list of secret keys
func getSecretKeys(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	val, err := usedSecret.SecretKeys()
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultJSON, err := json.Marshal(val)

	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	TResult.Print(string(resultJSON))
	w.Header().Set("Content-Type", MimeTypeJSON)
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

// saveSecret save the the secret key value received as parameter
func saveSecret(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var req StorePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = usedSecret.StoreSecretString(req.Key, req.Value)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	TResult.Printf(VALUE_STORED, req.Key, req.Value)
	w.WriteHeader(http.StatusCreated)
}

// deleteSecret deletes the secret content corresponding to the key received
// as parameter
func deleteSecret(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	vars := mux.Vars(r)
	id := vars["id"]

	if FilterSecretKeyFound(id, id, w) {
		return
	}

	b, err := usedSecret.DeleteSecret(id)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "Deleting a key", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if b {
		TResult.Printf(VALUE_DELETED, id)
	}
	w.WriteHeader(http.StatusOK)
}
