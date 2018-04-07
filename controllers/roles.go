package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/api-collection/repositories"
	"github.com/gorilla/mux"
	"strconv"
)

// RoleIndex define the logic for the routes GET /roles
func RoleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	roles := repositories.RolesGetAll()

	json.NewEncoder(w).Encode(roles)

}

// RoleCreate define the logic for the routes POST /roles
func RoleCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var role models.Role

	err = json.Unmarshal(body, &role)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			log.Error(err.Error())
		}
	} else {
		err = repositories.RoleCreate(&role)
		if err != nil {
			err := models.JSONError{Message: err.Error(), Code: 403}
			json.NewEncoder(w).Encode(err)
		} else {
			if role.ID != 0 {
				json.NewEncoder(w).Encode(role)
			} else {
				err := models.JSONError{Message: "Error when creating model", Code: 422}
				json.NewEncoder(w).Encode(err)
			}
		}
	}
}

// RoleShow contains the logic to display a specific role
func RoleShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
	}

	role := repositories.RoleGetByID(id)
	if role.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "Role Not Found", Code: 404})
	} else {
		json.NewEncoder(w).Encode(role)
	}
}

func RoleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err.Error())
	}

	role := repositories.RoleGetByID(id)
	if role.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "Role Not Found", Code: 404})
	} else {
		err = json.Unmarshal(body, &role)
		if err != nil {
			log.Error(err.Error())
		}

		err = repositories.RoleUpdate(role)
		if err != nil {
			json.NewEncoder(w).Encode(models.JSONError{Message: err.Error(), Code: 403})
		} else {
			json.NewEncoder(w).Encode(role)
		}
	}
}

func RoleDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
	}

	role := repositories.RoleGetByID(id)
	if role.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "Role Not Found", Code: 404})
	} else {
		err = repositories.RoleDelete(role)
		if err != nil {
			json.NewEncoder(w).Encode(models.JSONError{Message: err.Error(), Code: 403})
		} else {
			json.NewEncoder(w).Encode(models.JSONError{Message: "Role successfully deleted", Code: 200})
		}
	}
}
