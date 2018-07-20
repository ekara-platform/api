package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetInfo(t *testing.T) {
	logger = *log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	initLog(false, false)
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
