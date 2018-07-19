package api

const (
	API_CALLED_METHOD string = "API call %v, method %v"
	API_CALLED_PARAM  string = "API call with param %s = %v"
	API_CALLED_BODY   string = "API call with body = %v"

	API_RUNNING_ON string = "API listening on port %s"

	FLAGGED_WITH string = "flaged with %s : %v"

	ENVIRONMENT_CREATED string = "Environment %s created"

	USER_LOGIN          string = "User %s logged, token %s"
	USER_LOGOUT         string = "User logged out, token %s"
	USER_LOGIN_REQUIRED string = "Login required"

	LOGGER_INITIALIZED string = "logger initialized"
	TIME_REPORT        string = "execution of (%s:%d) took %s \n"

	ERROR_NO_BODY string = "Please send a request body"
)
