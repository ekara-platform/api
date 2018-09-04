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

	contentType := r.Header.Get("Content-type")
	s := getStorage()

	var b bool
	var result []byte
	var err error

	if contentType == MimeTypeYAML {
		b, result, err = s.Get(KEY_STORE_ENVIRONMENT_YAML_CONTENT)
		contentType = MimeTypeYAML
	} else {
		b, result, err = s.Get(KEY_STORE_ENVIRONMENT_JSON_CONTENT)
		contentType = MimeTypeJSON
	}

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

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func validate(e EnvironmentLoadRequest) (lagoon engine.Lagoon, err error, vErrs model.ValidationErrors) {
	root, flavor := repositoryFlavor(e.Location)
	s := getStorage()
	b, err := s.Contains(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
	if err != nil {
		// TODO make a proper error here
		return
	}

	var lagoonError error
	if b {
		var val []byte
		_, val, err = s.Get(KEY_STORE_ENVIRONMENT_PARAM_CONTENT)
		if err != nil {
			// TODO make a proper error here
			return
		}
		var p engine.ParamContent
		p, err = parseParams(string(val))
		if err != nil {
			// TODO make a proper error here
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
		var ok bool
		vErrs, ok = lagoonError.(model.ValidationErrors)
		// if the error is not a "validation error" then we return it
		if !ok {
			// TODO make a proper error here
			err = vErrs
			return
		} else {
			// We are facing validation errors
			err = nil
			return
		}
	}
	return
}

// checkEnvironment validate the enviroment received as parameter
func checkEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var e EnvironmentLoadRequest
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err, vErrs := validate(e)
	// An error occured validating the enviroment
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// The validation results into validation errors
	if vErrs.HasErrors() || vErrs.HasWarnings() {
		b, e := vErrs.JSonContent()
		if e != nil {
			http.Error(
				w,
				fmt.Errorf(ERROR_CONTENT, "marshalling the environment validation Json content:", err.Error()).Error(),
				http.StatusInternalServerError,
			)
			return
		}
		// Return both errors and warnings
		TLog.Printf(string(b))
		w.Header().Set("Content-Type", MimeTypeJSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		return
	} else {
		// The environment validation is okay
		TLog.Printf("Validation ok")
		w.WriteHeader(http.StatusOK)
		return
	}
}

// updateEnvironment updates the enviroment received as parameter
func updateEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var e EnvironmentLoadRequest
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lagoon, err, vErrs := validate(e)
	// An error occured validating the enviroment
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// The validation results into validation errors
	if vErrs.HasErrors() || vErrs.HasWarnings() {
		b, e := vErrs.JSonContent()
		if e != nil {
			http.Error(
				w,
				fmt.Errorf(ERROR_CONTENT, "marshalling the environment validation Json content:", err.Error()).Error(),
				http.StatusInternalServerError,
			)
		}
		// Return both errors and warnings
		TLog.Printf(string(b))
		w.Header().Set("Content-Type", MimeTypeJSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		return
	}
	// The environment validation is okay
	TLog.Printf("Validation ok")

	// TODO Complete the environment update process here

	environmentJson, err := json.Marshal(lagoon.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Json content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	environmentYaml, err := yaml.Marshal(lagoon.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Yaml content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	s := getStorage()
	if b := storeEnvironmentJSONContent(s, w, environmentJson); b {
		return
	}

	if b := storeEnvironmentYAMLContent(s, w, environmentYaml); b {
		return
	}

	if b := storeEnvironmentLocation(s, w, e.Location); b {
		return
	}
	storeEnvironmentTime(s, w)

	TResult.Printf(ENVIRONMENT_UPDATED, e.Location)
	w.WriteHeader(http.StatusCreated)
	return
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

	lagoon, err, vErrs := validate(e)
	// An error occured validating the enviroment
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// The validation results into validation errors
	if vErrs.HasErrors() || vErrs.HasWarnings() {
		b, e := vErrs.JSonContent()
		if e != nil {
			http.Error(
				w,
				fmt.Errorf(ERROR_CONTENT, "marshalling the environment validation Json content:", err.Error()).Error(),
				http.StatusInternalServerError,
			)
		}
		// Return both errors and warnings
		TLog.Printf(string(b))
		w.Header().Set("Content-Type", MimeTypeJSON)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		return
	}

	// The environment validation is okay
	TLog.Printf("Validation ok")

	// TODO Complete the environment creation process here

	environmentJson, err := json.Marshal(lagoon.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Json content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	environmentYaml, err := yaml.Marshal(lagoon.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Yaml content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	s := getStorage()
	if b := storeEnvironmentJSONContent(s, w, environmentJson); b {
		return
	}

	if b := storeEnvironmentYAMLContent(s, w, environmentYaml); b {
		return
	}

	if b := storeEnvironmentLocation(s, w, e.Location); b {
		return
	}

	storeEnvironmentTime(s, w)

	TResult.Printf(ENVIRONMENT_CREATED, e.Location)
	w.WriteHeader(http.StatusCreated)
	return
}

func storeEnvironmentYAMLContent(s Storage, w http.ResponseWriter, content []byte) (shouldReturn bool) {
	err := s.Store(KEY_STORE_ENVIRONMENT_YAML_CONTENT, content)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Storing the environment YAML content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		shouldReturn = true
	}
	return
}

func storeEnvironmentJSONContent(s Storage, w http.ResponseWriter, content []byte) (shouldReturn bool) {
	err := s.Store(KEY_STORE_ENVIRONMENT_JSON_CONTENT, content)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Storing the environment JSON content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		shouldReturn = true
	}
	return
}

func storeEnvironmentLocation(s Storage, w http.ResponseWriter, location string) (shouldReturn bool) {
	err := s.StoreString(KEY_STORE_ENVIRONMENT_LOCATION, location)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Storing the environment location:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		shouldReturn = true
	}
	return
}

func storeEnvironmentTime(s Storage, w http.ResponseWriter) (shouldReturn bool) {
	t := time.Now()
	err := s.StoreString(KEY_STORE_ENVIRONMENT_UPLOAD_TIME, t.Format("20060102150405"))
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Storing the environment upload time:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		shouldReturn = true
	}
	return
}

// deleteEnvironment deletes the enviroment currently
// manage by lagoon
func deleteEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	s := getStorage()
	_, err := s.Delete(KEY_STORE_ENVIRONMENT_JSON_CONTENT)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Deleting the environment Json content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	_, err = s.Delete(KEY_STORE_ENVIRONMENT_LOCATION)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Deleting the environment location:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	_, err = s.Delete(KEY_STORE_ENVIRONMENT_UPLOAD_TIME)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Deleting the environment upload time:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
	return
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
		return r, err
	}
	return r, nil
}
