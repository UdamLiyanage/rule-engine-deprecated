package main

type Rule struct {
	Key     string              `json:"_key"`
	Serial  string              `json:"serial"`
	Rules   []map[string]string `json:"rules"`
	Type    string              `json:"action_type"`
	Action  map[string]string   `json:"action"`
	States  State               `json:"states,omitempty"`
	Payload map[string]string   `json:"payload,omitempty"`
}

type State struct {
	Current  string   `json:"current"`
	Possible []string `json:"possible"`
}
