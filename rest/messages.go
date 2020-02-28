package rest

const (
	API_CALLED_METHOD string = "API call %v, method %v"
	API_CALLED_PARAM  string = "API call with param %s = %v"
	API_CALLED_BODY   string = "API call with body = %v"

	ENVIRONMENT_CREATED string = "Environment %s created"
	ENVIRONMENT_UPDATED string = "Environment %s updated"

	USER_LOGIN  string = "User %s logged, token %s"
	USER_LOGOUT string = "User logged out, token %s"
	TIME_REPORT string = "execution of (%s:%d) took %s \n"

	ERROR_NO_BODY string = "Please send a request body"
	ERROR_CONTENT string = "An error occured %s: %s"

	VALUE_STORED  string = "Value %s:%s has been stored"
	VALUE_DELETED string = "Value %s has been deleted"
)
