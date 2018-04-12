package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/api-collection/repositories"
	"io/ioutil"
)

// UserIndex define the logic for the routes GET /users
func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	users := repositories.UsersGetAll()

	json.NewEncoder(w).Encode(users)

}

// UserCreate define the logic for the routes POST /users
func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			log.Error(err.Error())
		}
	} else {
		user.Password = HashPassword(user.Password)

		err = repositories.UserCreate(&user)
		if err != nil {
			err := models.JSONError{Message: err.Error(), Code: 422}
			json.NewEncoder(w).Encode(err)
		} else {
			if user.ID != 0 {
				json.NewEncoder(w).Encode(user)
			} else {
				err := models.JSONError{Message: "Error when creating user", Code: 422}
				json.NewEncoder(w).Encode(err)
			}
		}
	}
}

// UserShow contains the logic to display a specific role
func UserShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
	}

	user := repositories.UserGetByID(id)
	if user.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "User Not Found", Code: 404})
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func UserGetRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Error(err)
	}

	user := repositories.UserGetByID(id)
	if user.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "User Not Found", Code: 404})
	} else {
		role := repositories.UserGetRole(user)
		if role.ID == 0 {
			json.NewEncoder(w).Encode(models.JSONError{Message: "User Not Found", Code: 404})
		} else {
			json.NewEncoder(w).Encode(role)
		}
	}
}