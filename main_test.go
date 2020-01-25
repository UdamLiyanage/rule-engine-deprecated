package main

import (
	"net/http"
	"testing"
)

func TestMQTTRoute(t *testing.T) {
	testRequestStatusCode("POST", "/mqtt-broker", nil, http.StatusOK, t)
}

func TestInvalidRoute(t *testing.T) {
	testRequestStatusCode("POST", "/mqtt", nil, http.StatusNotFound, t)
}
