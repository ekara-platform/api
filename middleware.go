package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Middleware func(http.Handler) http.Handler

type Middlewares []Middleware

func chain(middleware ...Middleware) Middlewares {
	return reverse(middleware)
}

func reverse(m Middlewares) Middlewares {
	for i := 0; i < len(m)/2; i++ {
		j := len(m) - i - 1
		m[i], m[j] = m[j], m[i]
	}
	return m
}

// The LogMw middleware logs the url and method for the http request
func LogMw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		TLog.Printf(API_CALLED_METHOD, r.URL, r.Method)
		h.ServeHTTP(w, r)
	})
}

// The BodyMw middleware checks that the request body is not empty
// and log its content.
// Its will return an http.StatusBadRequest if the request doesn't have a body
func BodyMw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := strings.TrimSpace(readReqContent(r))
		if r.Body == nil || s == "" {
			http.Error(w, ERROR_NO_BODY, http.StatusBadRequest)
			return
		}
		TLog.Printf(API_CALLED_BODY, s)
		h.ServeHTTP(w, r)
	})
}

// createParamMv allows to create a middleware in charge of log a specific
// path parameter value
func createParamMv(paramName string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			p := vars[paramName]
			TLog.Printf(API_CALLED_PARAM, paramName, p)
			h.ServeHTTP(w, r)
		})
	}
}

// The LoginMw middleware checks if the received token is valid or not
func LoginMw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getToken(r)
		// TODO Make a real test on the token valididty
		if token == "1234" {
			increaseTokenValidity(token)
		} else {
			http.Error(w, USER_LOGIN_REQUIRED, http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}
