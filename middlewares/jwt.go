package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/youkoulayley/api-collection/controllers"
	"github.com/youkoulayley/api-collection/models"
	"net/http"
	"strings"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}

					return controllers.JwtSalt, nil
				})
				if err != nil {
					json.NewEncoder(w).Encode(models.JSONError{Message: err.Error(), Code: 404})
					return
				}
				if token.Valid {
					context.Set(r, "decoded", token.Claims)
					next(w, r)
				} else {
					json.NewEncoder(w).Encode(models.JSONError{Message: "Invalid authorization token", Code: 403})
				}
			}
		} else {
			json.NewEncoder(w).Encode(models.JSONError{Message: "An authorization header is required", Code: 403})
		}
	})
}
