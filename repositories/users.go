package repositories

import (
	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
)

// UserGetByUsername fetch a user by its username
func UserGetByUsername(username string) *models.User {
	var user models.User

	bootstrap.Db().Where("name = ?", username).First(&user)

	return &user
}
