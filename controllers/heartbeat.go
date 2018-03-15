package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/youkoulayley/api-collection/models"
)

// HeartbeatIndex return the status of the api
func HeartbeatIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Heartbeat{Status: "OK", Code: "200"})
}
