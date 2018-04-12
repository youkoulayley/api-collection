package controllers

import (
	"golang.org/x/crypto/bcrypt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/youkoulayley/api-collection/repositories"
	"github.com/youkoulayley/api-collection/models"

	"github.com/dgrijalva/jwt-go"
	"time"
)

// Token struct contain the format for getting token
type Token struct {
	Username string
	Password string
}

type TokenString struct {
	Token 	string
}

var JwtSalt []byte

// HashPassword use the bcrypt generateFromPassword function but take a string parameters instead of a byte table.
func HashPassword(password string) string {
	bytePassword := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
	}

	return string(hashedPassword)
}

// CompareHashPassword use the bcrypt CompareHashAndPassword function but take strings in parameters instead of bytes
func CompareHashPassword(hash string, password string) bool {
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
		compare := CompareHashPassword(user.Password, tokenGet.Password)
		if compare {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)

			claims["username"] = user.Username
			claims["role"] = user.RoleID
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			tokenString, err := token.SignedString(JwtSalt)
			if err != nil {
				log.Info(err)
			}
			json.NewEncoder(w).Encode(TokenString{Token: tokenString})
		} else {
			json.NewEncoder(w).Encode(models.JSONError{Message: "Your login / Password is wrong", Code: 403})
		}
	}
}
