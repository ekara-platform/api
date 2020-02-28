package rest

import (
	"encoding/json"
	"net/http"

	_ "github.com/ekara-platform/engine"
	_ "github.com/ekara-platform/model"
)

type FakeInitRequest struct {
	Location string `json:"location"`
}

// fakeInit runs a fake initialization of the running api
func fakeInit(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var e FakeInitRequest
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO Read here the desciptor content

	TResult.Printf("Fake initialization done : %s", e.Location)
	w.WriteHeader(http.StatusCreated)
}
