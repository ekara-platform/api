package rest

import (
	"encoding/json"
	_ "fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// getValue returns the stored content corresponding to the key received
// as parameter
func getValue(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	vars := mux.Vars(r)
	id := vars["id"]
	if FilterKeyFound(id, id, w) {
		return
	}

	_, val, err := api.storageEngine.Get(id)
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

// getKeys returns the list of keys stored
func getKeys(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	val, err := api.storageEngine.Keys()
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

// saveValue save into the storage the key value received as parameter
func saveValue(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var req StorePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.storageEngine.StoreString(req.Key, req.Value)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	TResult.Printf(VALUE_STORED, req.Key, req.Value)
	w.WriteHeader(http.StatusCreated)
}

// deleteValue deletes the stored content corresponding to the key received
// as parameter
func deleteValue(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	vars := mux.Vars(r)
	id := vars["id"]

	if FilterKeyFound(id, id, w) {
		return
	}

	b, err := api.storageEngine.Delete(id)
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
