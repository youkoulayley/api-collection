package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/api-collection/repositories"
)

// UserShow contains the logic to display a specific role
func UserShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
	}

	role := repositories.RoleGetByID(id)
	if role.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "User Not Found", Code: 404})
	} else {
		json.NewEncoder(w).Encode(role)
	}
}
