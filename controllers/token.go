package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/api-collection/repositories"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

// Token struct contain the format for getting token
type Token struct {
	Username string
	Password []byte
}

// TokenGet has the logic to generate a token for user
func TokenGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var tokenGet Token

	err = json.Unmarshal(body, &tokenGet)
	if err != nil {
		log.Error(err)
	}

	// Verify that the user exists in database
	user := repositories.UserGetByUsername(tokenGet.Username)

	if user.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "User Not Found", Code: 404})
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword(tokenGet.Password, bcrypt.DefaultCost)
		if err != nil {
			log.Info(err)
		}
		fmt.Println(string(hashedPassword))

		if user.Username == tokenGet.Username && user.Password == string(hashedPassword) {
			json.NewEncoder(w).Encode(models.JSONError{Message: "Authorized", Code: 202})
		} else {
			json.NewEncoder(w).Encode(models.JSONError{Message: "Your login / Password is wrong", Code: 404})
		}
	}
}
