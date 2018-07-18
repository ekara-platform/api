package api

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middlewares Middlewares
}

func NewRouter(routes Routes) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(use(http.HandlerFunc(route.HandlerFunc), route.Middlewares))
	}
	return r
}

func use(h http.Handler, middleware Middlewares) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func readReqContent(r *http.Request) string {
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	return string(bodyBytes)
}
