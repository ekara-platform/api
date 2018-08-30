package api

import (
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(""))
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
		TLog.Printf("Creating lagoon environment with parameter for templating")
		lagoon, lagoonError = engine.Create(TLog, "/var/lib/lagoon", p)
	} else {
		TLog.Printf("Creating lagoon environment without parameter for templating")
		lagoon, lagoonError = engine.Create(TLog, "/var/lib/lagoon", map[string]interface{}{})
	}

	if lagoonError == nil {
		panic(lagoonError)
		lagoonError = lagoon.Init(root, flavor) // FIXME: really need custom descriptor name ?
	}

	if lagoonError != nil {
		vErrs, ok := lagoonError.(model.ValidationErrors)
		// if the error is not a "validation error" then we return it
		if !ok {
			// TODO make a proper error here
			http.Error(w, lagoonError.Error(), http.StatusInternalServerError)
		} else {
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
	w.Header().Set("Content-Type", "application/json")

	// TODO implement and change returned status

	w.WriteHeader(http.StatusNotImplemented)
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
