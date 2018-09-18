package api

import (
	"encoding/json"
	"strings"

	"github.com/gorilla/mux"

	"github.com/lagoon-platform/api/docker"

	"net/http"
	"net/url"
	"os"
)

const (
	STR string = "C:\\Users\\e518546\\.docker\\machine\\certs"
)

type ApiInfo struct {
	Version            string
	Host               string
	Url                string
	EnvironmentDetails *EnvironmentDetails `json:",omitempty"`
	Err                string
}

type EnvironmentDetails struct {
	url                string
	storedShortContent map[string]string
	storedLongContent  map[string]string
}

func (r EnvironmentDetails) MarshalJSON() ([]byte, error) {
	t := struct {
		Content map[string]string
	}{
		Content: make(map[string]string),
	}

	for k, v := range r.storedShortContent {
		t.Content[k] = v
	}
	for k, v := range r.storedLongContent {
		t.Content[k] = v
	}
	return json.Marshal(t)
}

func (d *EnvironmentDetails) addStoredShortContent(s Storage, key string) {
	b, val, err := s.Get(key)
	if err != nil {
		d.storedShortContent[key] = err.Error()
		return
	}
	if b {
		d.storedShortContent[key] = string(val)
	}
}

func (d *EnvironmentDetails) addStoredLongContent(s Storage, key string, route string) {
	if b, _ := s.Contains(key); b {
		r := application.Router.Get(route)

		template, err := r.GetPathTemplate()
		if err != nil {
			panic(err)
		}
		nr := mux.NewRouter()
		nr.Host(d.url).
			Path(template).
			Name(route)

		var url *url.URL
		var e error
		// TODO parse this to get the named param...

		if strings.Contains(template, "{id}") {
			url, e = nr.Get(route).URL("id", key)
			if e != nil {
				panic(e)
			}
		} else {
			url, e = nr.Get(route).URL()
			if e != nil {
				panic(e)
			}
		}
		d.storedLongContent[removeLagoonPrefix(key)] = string(url.String())
	}
}

// getInfo returns the application detailed informations
func getInfo(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	TLog.Println("Calling The docker stuff...")
	docker.TestDocker(TLog, "tcp://192.168.99.100:2376", "1.30", STR)

	w.Header().Set("Content-Type", MimeTypeJSON)

	s := getStorage()

	info := ApiInfo{
		Version: application.Version,
		Url:     r.Host,
		EnvironmentDetails: &EnvironmentDetails{
			url:                r.Host,
			storedShortContent: make(map[string]string),
			storedLongContent:  make(map[string]string),
		},
	}

	name, err := os.Hostname()
	if err != nil {
		info.Host = err.Error()
	}
	info.Host = name
	d := info.EnvironmentDetails
	d.addStoredShortContent(s, KEY_STORE_ENV_LOCATION)
	d.addStoredShortContent(s, KEY_STORE_ENV_CREATED_AT)
	d.addStoredShortContent(s, KEY_STORE_ENV_UPDATED_AT)

	d.addStoredLongContent(s, KEY_STORE_ENV_PARAM, "GetEnvironmentParam")
	d.addStoredLongContent(s, KEY_STORE_ENV_JSON, "GetStorage")
	d.addStoredLongContent(s, KEY_STORE_ENV_YAML, "GetStorage")

	infoJSON, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	TResult.Print(string(infoJSON))
	w.Write(infoJSON)
}
