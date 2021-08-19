package models

type ResponseWatson struct {
	Context interface{} `json:"context"`
	Output  interface{} `json:"output"`
}
