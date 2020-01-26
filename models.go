package main

type MQTTPublish struct {
	Broker string `json:"broker_type"`
	URL    string `json:"broker_url"`
	Topic  string `json:"topic"`
}

type HTTPPublish struct {
	Type string `json:"request_type"`
	URL  string `json:"url"`
}

type Rule struct {
	Serial  string              `json:"serial"`
	Rules   []map[string]string `json:"rules"`
	Type    string              `json:"action_type"`
	Action  map[string]string   `json:"action"`
	Payload map[string]string   `json:"payload,omitempty"`
}
