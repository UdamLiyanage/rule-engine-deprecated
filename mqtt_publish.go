package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func publish(config, payload map[string]string, states State, trigger bool, key string) {
	fmt.Println("MQTT Publish")
	var reqBody []byte
	opts := MQTT.NewClientOptions().AddBroker(config["broker_url"])
	opts.SetClientID("platform")

	if trigger {
		reqBody = triggerDevice(states, payload, key)
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish(config["topic"], 0, false, reqBody)
	res := token.WaitTimeout(2 * time.Second)
	if res {
		fmt.Println("Success")
	} else {
		fmt.Println("Request Failed")
	}
	c.Disconnect(250)
}
