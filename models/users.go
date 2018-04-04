package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username	string	`gorm:"type:varchar(20);unique_index"`
	Password	string	`gorm:"type:varchar(255)"`
	RoleID		int
	Role		Role
}