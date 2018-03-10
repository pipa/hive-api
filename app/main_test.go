package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestNewApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app.Application)
	testKey := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InJrWC1tc1VzVyIsImFjY2VzcyI6WyJ3ZWIiXSwiaWF0IjoxNTA2MzU1NTcxLCJpc3MiOiJSRXgtQVBJQGRldiJ9.gzGS2SXemPUf_-3mO6rfpQ-2OydeX9ch_9ZuqJL5pyM"

	// Not Auth
	e.GET("/ip/127.0.0.1").Expect().Status(httptest.StatusUnauthorized)

	// Get info for localhost ip
	e.GET("/ip/127.0.0.1").WithHeader("Authorization", testKey).Expect().Status(httptest.StatusOK)

	// test not found
	e.GET("/notfound").Expect().Status(httptest.StatusNotFound)
}
