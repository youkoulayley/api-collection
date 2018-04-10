package repositories

import (
	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
	"errors"
	log "github.com/sirupsen/logrus"
)

// UsersGetAll fetch all users
func UsersGetAll() []models.User {
	var u []models.User

	bootstrap.Db().Find(&u)

	return u
}

// UserCreate create a new role
func UserCreate(u *models.User) error {
	if u == nil {
		log.Error(u)
	}

	if u.Username != "" && u.Email != "" && u.Password != "" {
		// Check if role already exists
		var count uint
		bootstrap.Db().Where("username = ? AND email = ?", u.Username, u.Email).Find(&models.User{}).Count(&count)

		if count == 0 {
			err := bootstrap.Db().Create(&u).Error
			if err != nil {
				log.Error(err)
				return errors.New(err.Error())
			}
		}
	} else {
		return errors.New("user with empty name, email and password is forbidden")
	}
	return nil
}

// UserGetByID get a user by its ID
func UserGetByID(id int) *models.User {
	var u models.User

	bootstrap.Db().First(&u, id)

	return &u
}

// UserGetByUsername fetch a user by its username
func UserGetByUsername(username string) *models.User {
	var user models.User

	bootstrap.Db().Where("name = ?", username).First(&user)

	return &user
}
