package models

// Heartbeat is the models for the heartbeat model
type Heartbeat struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
