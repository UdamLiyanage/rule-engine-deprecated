package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/gin-gonic/gin"
	"strconv"
)

func inferRules(c *gin.Context) {
	var payload map[string]string
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		panic(err)
	}
	rules := getRules("test_serial")
	for _, rule := range rules {
		if compareRules(rule.Rules, payload) {
			executeAction(rule.Type, rule.Action)
		}
	}
	c.Status(200)
}

func getRules(serial string) []Rule {
	var res []Rule
	bindVars := map[string]interface{}{
		"serial": serial,
	}
	cursor, err := database.Query(context.TODO(), "FOR r in rules FILTER r.serial==@serial RETURN r", bindVars)
	if err != nil {
		panic(err)
	}
	for {
		var doc Rule
		_, err := cursor.ReadDocument(context.TODO(), &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		res = append(res, doc)
	}
	if res == nil {
		panic(err)
	}
	return res
}

func compareRules(rules []map[string]string, payload map[string]string) bool {
	match := false
	for _, compare := range rules {
		fmt.Println(compare["parameter"] + compare["condition"])
		if _, ok := payload[compare["parameter"]]; ok {
			value, err := strconv.Atoi(compare["value"])
			if err != nil {
				panic(err)
			}
			payloadValue, err := strconv.Atoi(payload[compare["parameter"]])
			if err != nil {
				panic(err)
			}
			switch compare["condition"] {
			case "equals":
				if value == payloadValue {
					match = true
				}
			case "greaterequals":
				if value >= payloadValue {
					match = true
				}
			case "greater":
				if value > payloadValue {
					match = true
				}
			case "lessequals":
				if value <= payloadValue {
					match = true
				}
			case "less":
				if value < payloadValue {
					match = true
				}
			default:
				break
			}
		} else {
			println("Wrong Payload!")
		}
	}
	return match
}

func executeAction(actionType string, action map[string]string) {
	switch actionType {
	case "mqtt_publish":
		fmt.Println("MQTT Publish")
	case "url_call":
		fmt.Println("URL Call")
	default:
		fmt.Println("Log to DB")
	}
}
