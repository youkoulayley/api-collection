package controllers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/api-collection/repositories"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"time"
)

type auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type token struct {
	Token string `json:"token"`
}

// JwtSalt is getting its value from main.go
var JwtSalt []byte

// hashPassword use the b-crypt generateFromPassword function but take a string parameters instead of a byte table.
func hashPassword(password string) string {
	bytePassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
	}

	return string(hashedPassword)
}

// CompareHashPassword use the b-crypt CompareHashAndPassword function but take strings in parameters instead of bytes
func compareHashPassword(hash string, password string) bool {
	hashByte := []byte(hash)
	passwordByte := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashByte, passwordByte)
	if err != nil {
		return false
	}
	return true
}

// TokenGet has the logic to generate a token for user
func TokenGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var a auth

	err = json.Unmarshal(body, &a)
	if err != nil {
		log.Error(err)
	}

	// Verify that the user exists in database
	user := repositories.UserGetByUsername(a.Username)

	if user.ID == 0 {
		json.NewEncoder(w).Encode(models.JSONError{Message: "User Not Found", Code: 404})
	} else {
		compare := compareHashPassword(user.Password, a.Password)
		if compare {
			jwtToken := jwt.New(jwt.SigningMethodHS256)
			claims := jwtToken.Claims.(jwt.MapClaims)

			claims["username"] = user.Username
			claims["role"] = user.RoleID
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			tokenString, errSign := jwtToken.SignedString(JwtSalt)
			if errSign != nil {
				log.Info(err)
			}
			json.NewEncoder(w).Encode(token{Token: tokenString})
		} else {
			json.NewEncoder(w).Encode(models.JSONError{Message: "Your login / Password is wrong", Code: 403})
		}
	}
}
