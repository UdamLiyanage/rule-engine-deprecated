package main

import (
	"net/http"
	"testing"
)

func TestInvalidRoute(t *testing.T) {
	testRequestStatusCode("POST", "/mqtt", nil, http.StatusNotFound, t)
}
