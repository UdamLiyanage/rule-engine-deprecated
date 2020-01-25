package main

import "testing"

func TestInferRules(t *testing.T) {
	testRequestStatusCode("POST", "/mqtt-broker", nil, 200, t)
}
