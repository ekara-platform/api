package main

import (
	"net/http"
	"testing"
)

func TestGetInfo(t *testing.T) {
	req, _ := http.NewRequest("GET", "/infos/", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp)
	if body := getNotEmptyBody(t, resp); body != "" {
		if b, result := checkJsonRoundTpip(t, body, &ApiInfo{}); !b {
			t.Errorf(`Not the same received "%s" converted to "%s"`, body, result)
		}
	}
}
