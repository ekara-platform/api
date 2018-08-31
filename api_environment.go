package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lagoon-platform/engine"
	"github.com/lagoon-platform/model"

	"gopkg.in/yaml.v2"
)

type EnvironmentLoadRequest struct {
	Location string `json:"location"`
}

// getEnvironment returns the details enviroment currently
// manage by lagoon
func getEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	s := getStorage()
	b, resultJSON, err := s.Get(KEY_STORE_ENVIRONMENT_JSON_CONTENT)
	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !b {
		err := fmt.Errorf("The environment cannot be found into the storage")
		TLog.Printf(ERROR_CONTENT, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
	w.WriteHeader(http.StatusOK)
}

// loadEnvironment loads the enviroment received as parameter
func loadEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var e EnvironmentLoadRequest
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO Parse the file content here
	s := getStorage()
	root, flavor := repositoryFlavor(e.Location)
	b, err := s.Contains(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)

	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var lagoonError error
	var lagoon engine.Lagoon

	if b {
		_, val, err := s.Get(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
		if err != nil {
			// TODO make a proper error here
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		p, err := parseParams(string(val))
		if err != nil {
			// TODO make a proper error here
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		TLog.Printf("Creating lagoon environment with parameter for templating")
		lagoon, lagoonError = engine.Create(TLog, "/var/lib/lagoon", p)
	} else {
		TLog.Printf("Creating lagoon environment without parameter for templating")
		lagoon, lagoonError = engine.Create(TLog, "/var/lib/lagoon", map[string]interface{}{})
	}

	if lagoonError == nil {
		lagoonError = lagoon.Init(root, flavor) // FIXME: really need custom descriptor name ?
	}

	if lagoonError != nil {
		vErrs, ok := lagoonError.(model.ValidationErrors)
		// if the error is not a "validation error" then we return it
		if !ok {
			// TODO make a proper error here
			http.Error(w, lagoonError.Error(), http.StatusInternalServerError)
		} else {
			// We are facing validation errors
			TLog.Printf(lagoonError.Error())
			b, e := vErrs.JSonContent()
			if e != nil {
				// TODO make a proper error here
				http.Error(w, e.Error(), http.StatusInternalServerError)
			}
			// print both errors and warnings into the report file
			TLog.Printf(string(b))

		}
	} else {
		TLog.Printf("Validation ok")
	}

	environmentJson, err := json.Marshal(lagoon.Environment())
	if err != nil {
		// TODO make a proper error here
		panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.Store(KEY_STORE_ENVIRONMENT_JSON_CONTENT, environmentJson)
	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.StoreString(KEY_STORE_ENVIRONMENT_LOCATION, e.Location)
	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := time.Now()
	err = s.StoreString(KEY_STORE_ENVIRONMENT_UPLOAD_TIME, t.Format("20060102150405"))
	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	TResult.Printf(ENVIRONMENT_CREATED, e.Location)
	w.WriteHeader(http.StatusCreated)
}

// deleteEnvironment deletes the enviroment currently
// manage by lagoon
func deleteEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	s := getStorage()
	_, err := s.Delete(KEY_STORE_ENVIRONMENT_JSON_CONTENT)
	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = s.Delete(KEY_STORE_ENVIRONMENT_LOCATION)
	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = s.Delete(KEY_STORE_ENVIRONMENT_UPLOAD_TIME)
	if err != nil {
		// TODO make a proper error here
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

// TODO DUPLICATED FROM INSTALLER MOVE THIS INTO THE ENGINE
//repositoryFlavor returns the repository flavor, branchn tag ..., based on the
// presence of '@' into the given url
func repositoryFlavor(url string) (string, string) {

	if strings.Contains(url, "@") {
		s := strings.Split(url, "@")
		return s[0], s[1]
	}
	return url, ""
}

//ParseParams parses a yaml file into a map of "string:interface{}"
func parseParams(content string) (engine.ParamContent, error) {
	r := make(engine.ParamContent)

	err := yaml.Unmarshal([]byte(content), &r)
	if err != nil {
		panic(err)
	}
	return r, nil
}
