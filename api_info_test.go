package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetInfo(t *testing.T) {
	initLog(true, true)
	handler := http.HandlerFunc(getInfo)
	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusOK, resp)
	if body := getNotEmptyBody(t, resp); body != "" {
		if b, result := checkJsonRoundTpip(t, body, &ApiInfo{}); !b {
			t.Errorf(`Not the same received "%s" converted to "%s"`, body, result)
		}
	}
}
