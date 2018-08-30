package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type StorePostRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// getValue returns the stored content corresponding to the key received
// as parameter
func getValue(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	vars := mux.Vars(r)
	id := vars["id"]

	b, val, err := getStorage().Get(id)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", id)
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := StorePostRequest{
		id,
		string(val),
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	TResult.Print(string(resultJSON))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
	w.WriteHeader(http.StatusOK)
}

// saveValue save into the storage the key value received as parameter
func saveValue(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var req StorePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = getStorage().StoreString(req.Key, req.Value)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, err.Error())
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

	s := getStorage()

	b, err := s.Contains(id)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !b {
		err := fmt.Errorf("The key \"%s\" cannot be found", id)
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err = getStorage().Delete(id)
	if err != nil {
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if b {
		TResult.Printf(VALUE_DELETED, id)
	}
	w.WriteHeader(http.StatusOK)
}
