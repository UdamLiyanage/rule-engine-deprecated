package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func publish(config, payload map[string]string) {
	fmt.Println("MQTT Publish")
	opts := MQTT.NewClientOptions().AddBroker(config["broker_url"])
	opts.SetClientID("platform")

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish(config["topic"], 0, false, payload)
	res := token.WaitTimeout(2 * time.Second)
	if res {
		fmt.Println("Success")
	} else {
		fmt.Println("Request Failed")
	}
	c.Disconnect(250)
}
