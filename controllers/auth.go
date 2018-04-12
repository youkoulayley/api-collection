package controllers

import (
	"golang.org/x/crypto/bcrypt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/youkoulayley/api-collection/repositories"
	"github.com/youkoulayley/api-collection/models"
)

// Token struct contain the format for getting token
type Token struct {
	Username string
	Password string
}

func HashPassword(password string) string {
	bytePassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
	}

	return string(hashedPassword)
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
		err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(tokenGet.Password))
		log.Info(err)
		if err == nil {
			json.NewEncoder(w).Encode(models.JSONError{Message: "Authorized", Code: 202})
		} else {
			json.NewEncoder(w).Encode(models.JSONError{Message: "Your login / Password is wrong", Code: 403})
		}
	}
}
