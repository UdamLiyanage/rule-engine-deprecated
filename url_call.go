package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func urlCall(config, body map[string]string) {
	fmt.Println("Url Call")
	var err error
	var req *http.Request
	var payload []byte
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	if body != nil {
		payload, err = json.Marshal(body)
		if err != nil {
			panic(err)
		}
	}
	method := config["request_type"]
	switch method {
	case "POST":
		req, err = http.NewRequest("POST", config["url"], bytes.NewBuffer(payload))
	case "GET":
		req, err = http.NewRequest("GET", config["url"], nil)
	case "PUT":
		req, err = http.NewRequest("PUT", config["url"], bytes.NewBuffer(payload))
	case "DELETE":
		req, err = http.NewRequest("DELETE", config["url"], nil)
	default:
		break
	}
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
}
