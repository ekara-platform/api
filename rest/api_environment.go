package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ekara-platform/api/storage"
	"github.com/ekara-platform/engine"
	"github.com/ekara-platform/engine/ansible"
	"github.com/ekara-platform/model"

	"gopkg.in/yaml.v2"
)

type EnvironmentLoadRequest struct {
	Location string `json:"location"`
}

// getEnvironment returns the details enviroment currently
// manage by ekara
// Acceptable Content type are :
// - "application/json"
// - "application/x.yaml"
func getEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	if FilterKeyFound(storage.KEY_STORE_ENV_LOCATION, "Environment", w) {
		return
	}

	contentType := r.Header.Get("Content-type")
	s := api.storageEngine

	var b bool
	var result []byte
	var err error

	if contentType == MimeTypeYAML {
		TLog.Println("Getting YAML environment")
		b, result, err = s.Get(storage.KEY_STORE_ENV_YAML)
		contentType = MimeTypeYAML
	} else {
		// Returns JSON by default
		TLog.Println("Getting JSON environment")
		b, result, err = s.Get(storage.KEY_STORE_ENV_JSON)
		contentType = MimeTypeJSON
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !b {
		err := fmt.Errorf("The environment cannot be found into the storage")
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func validate(e EnvironmentLoadRequest) (ekara engine.Engine, err error, vErrs model.ValidationErrors) {
	root, flavor := repositoryFlavor(e.Location)
	s := api.storageEngine
	b, err := s.Contains(storage.KEY_STORE_ENV_PARAM)
	if err != nil {
		// TODO make a proper error here
		return
	}

	var ekaraError error
	if b {
		var val []byte
		_, val, err = s.Get(storage.KEY_STORE_ENV_PARAM)
		if err != nil {
			// TODO make a proper error here
			return
		}
		var p ansible.ParamContent
		p, err = parseParams(string(val))
		if err != nil {
			// TODO make a proper error here
			return
		}
		TLog.Printf("Creating ekara environment with parameter for templating")
		ekara, ekaraError = engine.Create(TLog, "/var/lib/ekara", p)
	} else {
		TLog.Printf("Creating ekara environment without parameter for templating")
		ekara, ekaraError = engine.Create(TLog, "/var/lib/ekara", map[string]interface{}{})
	}

	if ekaraError == nil {
		ekaraError = ekara.Init(root, flavor) // FIXME: really need custom descriptor name ?
	}

	if ekaraError != nil {
		var ok bool
		vErrs, ok = ekaraError.(model.ValidationErrors)
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

	if FilterKeyFound(storage.KEY_STORE_ENV_LOCATION, "Environment", w) {
		return
	}

	var e EnvironmentLoadRequest
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ekara, err, vErrs := validate(e)
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

	environmentJson, err := json.Marshal(ekara.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Json content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	environmentYaml, err := yaml.Marshal(ekara.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Yaml content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	s := api.storageEngine
	if b := storeEnvironmentJSONContent(s, w, environmentJson); b {
		return
	}

	if b := storeEnvironmentYAMLContent(s, w, environmentYaml); b {
		return
	}

	if b := storeEnvironmentLocation(s, w, e.Location); b {
		return
	}
	storeEnvironmentTime(storage.KEY_STORE_ENV_UPDATED_AT, s, w)

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

	s := api.storageEngine
	b, err := s.Contains(storage.KEY_STORE_ENV_LOCATION)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if b {
		err := fmt.Errorf("An environment has already been created")
		TLog.Printf(ERROR_CONTENT, "", err.Error())
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	ekara, err, vErrs := validate(e)
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
	}

	// The environment validation is okay
	TLog.Printf("Validation ok")

	// TODO Complete the environment creation process here

	environmentJson, err := json.Marshal(ekara.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Json content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	environmentYaml, err := yaml.Marshal(ekara.Environment())
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "marshalling the environment Yaml content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	if b := storeEnvironmentJSONContent(s, w, environmentJson); b {
		return
	}

	if b := storeEnvironmentYAMLContent(s, w, environmentYaml); b {
		return
	}

	if b := storeEnvironmentLocation(s, w, e.Location); b {
		return
	}

	if b := storeEnvironmentTime(storage.KEY_STORE_ENV_CREATED_AT, s, w); b {
		return
	}

	TResult.Printf(ENVIRONMENT_CREATED, e.Location)
	w.WriteHeader(http.StatusCreated)
	return
}

func storeEnvironmentYAMLContent(s storage.Storage, w http.ResponseWriter, content []byte) (shouldReturn bool) {
	shouldReturn = store(s, w, storage.KEY_STORE_ENV_YAML, content, "Storing the environment YAML content:")
	return
}

func storeEnvironmentJSONContent(s storage.Storage, w http.ResponseWriter, content []byte) (shouldReturn bool) {
	shouldReturn = store(s, w, storage.KEY_STORE_ENV_JSON, content, "Storing the environment JSON content:")
	return
}

func storeEnvironmentLocation(s storage.Storage, w http.ResponseWriter, location string) (shouldReturn bool) {
	shouldReturn = store(s, w, storage.KEY_STORE_ENV_LOCATION, []byte(location), "Storing the environment location:")
	return
}

func storeEnvironmentTime(key string, s storage.Storage, w http.ResponseWriter) (shouldReturn bool) {
	t := time.Now()
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	shouldReturn = store(s, w, key, []byte(t.Format(layout)), "Storing "+key)
	return
}

func store(s storage.Storage, w http.ResponseWriter, key string, content []byte, message string) (callerShouldReturn bool) {
	err := s.Store(key, content)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, message, err.Error()).Error(),
			http.StatusInternalServerError,
		)
		callerShouldReturn = true
	}
	return
}

// deleteEnvironment deletes the enviroment currently
// manage by ekara
func deleteEnvironment(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	if FilterKeyFound(storage.KEY_STORE_ENV_LOCATION, "Environment", w) {
		return
	}

	s := api.storageEngine
	_, err := s.Delete(storage.KEY_STORE_ENV_JSON)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Deleting the environment Json content:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	_, err = s.Delete(storage.KEY_STORE_ENV_LOCATION)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Deleting the environment location:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	_, err = s.Delete(storage.KEY_STORE_ENV_CREATED_AT)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Deleting the environment creation time:", err.Error()).Error(),
			http.StatusInternalServerError,
		)
		return
	}

	_, err = s.Delete(storage.KEY_STORE_ENV_UPDATED_AT)
	if err != nil {
		http.Error(
			w,
			fmt.Errorf(ERROR_CONTENT, "Deleting the environment update time:", err.Error()).Error(),
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
func parseParams(content string) (ansible.ParamContent, error) {
	r := make(ansible.ParamContent)
	err := yaml.Unmarshal([]byte(content), &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
