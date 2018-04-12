package controllers

import (
	"encoding/json"
	"net/http"
)

type heartbeat struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// HeartbeatIndex return the status of the api
func HeartbeatIndex(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(heartbeat{Status: "OK", Code: 200})
}
