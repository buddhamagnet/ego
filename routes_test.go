package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Package level router
var m *mux.Router

// Package level test HTTP response recorder.
var res *httptest.ResponseRecorder

// setup sets the environment variable containing the
// RAML file location, boots the router and fires up
// an HTTP response recorder for each test.
func setup() {
	os.Setenv("RAML_FILE_PATH", "ego.raml")
	m = NewRouter()
	res = httptest.NewRecorder()
}

// teardown can be used for cleanup after the tests have run
func teardown() {}

func WithContext(t *testing.T, test func(*testing.T)) {
	setup()
	test(t)
	teardown()
}

func TestRootHandler(t *testing.T) {
	WithContext(t, func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		// Call the endpoint.
		m.ServeHTTP(res, req)
		// Check status code from root endpoint.
		assert.Equal(t, res.Code, 200)
	})
}
