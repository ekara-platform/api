package api

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	User     string `json:"user"`
	Password string `json:"pwd"`
}

type LoginResponse struct {
	User     string `json:"user"`
	Password string `json:"pwd"`
}

// login performs a login request
func login(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	var l LoginRequest
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	TLog.Printf("request login for %s/%s", l.User, l.Password)
	if l.User == "nirekin" && l.Password == "1234" {
		// TODO Parse the struct content here
		// Log the user in and return the token
		token := createToken()
		// TODO return the token here
		r.Header.Set("lagoon_token", token)
		TResult.Printf(USER_LOGIN, l.User, token)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// logout performs a login request
func logout(w http.ResponseWriter, r *http.Request) {
	defer traceTime(here())()

	token := getToken(r)
	TLog.Printf("request logout for %s", token)
	deleteToken()

	TResult.Printf(USER_LOGOUT, token)
	w.WriteHeader(http.StatusNotImplemented)
}

func createToken() (token string) {
	// TODO Make real implementation here
	token = "123456789"
	return
}

func deleteToken() {
	// TODO Make real implementation here
}

func increaseTokenValidity(token string) {
	// TODO make real implementation here
	TLog.Printf("token %s validity increased", token)
}

func getToken(r *http.Request) (token string) {
	// TODO make real implementation here
	token = r.Header.Get("lagoon_token")
	TLog.Printf("get token :%s ", token)
	return
}
