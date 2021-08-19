package models

import assistant "github.com/watson-developer-cloud/go-sdk/v2/assistantv1"

type RequestWatson struct {
	Context *assistant.Context `json:"context"`
	Input   string             `json:"input"`
}
