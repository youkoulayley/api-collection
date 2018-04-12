package repositories

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
)

// RolesGetAll fetch all roles
func RolesGetAll() []models.Role {
	var r []models.Role

	bootstrap.Db().Find(&r)

	return r
}

// RoleCreate create a new role
func RoleCreate(r *models.Role) error {
	if r == nil {
		log.Error(r)
	}

	if r.Name != "" {
		// Check if role already exists
		var count uint
		bootstrap.Db().Where("name = ?", r.Name).Find(&models.Role{}).Count(&count)

		if count == 0 {
			err := bootstrap.Db().Create(&r).Error
			if err != nil {
				return errors.New(err.Error())
			}
		}
	} else {
		return errors.New("role with empty name if forbidden")
	}
	return nil
}

// RoleGetByID fetch a role by its ID
func RoleGetByID(id int) *models.Role {
	var role models.Role

	bootstrap.Db().First(&role, id)

	return &role
}

// RoleUpdate update a role
func RoleUpdate(r *models.Role) error {
	if r.Name != "" {
		// Check if name already exists
		var count uint
		bootstrap.Db().Where("name = ? AND id <> ?", r.Name, r.ID).Find(&models.Role{}).Count(&count)

		if count == 0 {
			err := bootstrap.Db().Save(r).Error
			if err != nil {
				return errors.New(err.Error())
			}
		} else {
			return errors.New("role with this name already exists")
		}
	} else {
		return errors.New("role with empty name if forbidden")
	}
	return nil
}

//RoleDelete delete a role by its ID
func RoleDelete(r *models.Role) error {
	err := bootstrap.Db().Delete(&r).Error
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
