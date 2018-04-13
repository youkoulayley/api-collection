package middlewares

import (
	"net/http"
	"github.com/youkoulayley/api-collection/controllers"
	"github.com/youkoulayley/api-collection/models"
	"encoding/json"
)

// IsAdmin check if a user is admin
func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JwtToken := controllers.DecodeJwtToken(r)

		if JwtToken.Role == 1 {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.JSONError{Message: "Unauthorized", Code: 403})
		}
	})
}
