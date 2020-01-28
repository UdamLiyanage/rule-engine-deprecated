package main

import (
	"context"
	"encoding/json"
)

func triggerDevice(states State, payload map[string]string, key string) []byte {
	current := states.Current
	trigger := ""
	for _, v := range states.Possible {
		if v != current {
			trigger = v
			break
		}
	}
	states.Current = trigger
	updateDatabase(&states, key)

	for k, _ := range payload {
		payload[k] = trigger
	}

	rPayload, err := json.Marshal(payload)
	if err != nil {
		println("Trigger Device Marshal Error 1")
	}
	return rPayload
}

func updateDatabase(reqBody *State, key string) {
	patch := map[string]interface{}{
		"states": &reqBody,
	}
	col, err := database.Collection(context.TODO(), "rules")
	if err != nil {
		panic(err)
	}
	_, err = col.UpdateDocument(context.TODO(), key, patch)
	if err != nil {
		panic(err)
	}
}
