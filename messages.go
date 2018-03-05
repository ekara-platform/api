package main

const (
	API_CALLED_METHOD string = "API call %v, method %v"
	API_CALLED_PARAM  string = "API call with param %s = %v"
	API_CALLED_BODY   string = "API call with body = %v"

	API_RUNNING_ON string = "API listening on port %s"

	FLAGGED_WITH string = "flaged with %s : %v"
	FLAG_SCRIPT  string = "log the api request as script"
	FLAG_TIME    string = "log the execution time"
	FLAG_PORT    string = "the running port of the app"

	ENVIRONMENT_CREATED string = "Environment %s created"
	LOGGER_INITIALIZED  string = "logger initialized"
	TIME_REPORT         string = "execution of (%s:%d) took %s \n"

	ERROR_NO_BODY string = "Please send a request body"
)