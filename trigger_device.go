package main

import (
	"context"
	"encoding/json"
)

func triggerDevice(states State, payload map[string]string, key string) []byte {
	current := states.Current
	trigger := ""
	length := len(states.Possible) - 1
	for k, v := range states.Possible {
		if v == current {
			if k+1 > length {
				trigger = states.Possible[0]
				break
			} else {
				trigger = states.Possible[k+1]
				break
			}
		}
	}
	states.Current = trigger
	updateDatabase(&states, key)

	for k := range payload {
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
