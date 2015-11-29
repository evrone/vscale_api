package vscale

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

const (
	testToken = "testtoken"
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = NewClient(testToken, server.URL)
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}

	if token := r.Header.Get("X-Token"); token != testToken {
		t.Errorf("X-Token should be %s, but is %s", testToken, token)
	}
}
