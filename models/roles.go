package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name		string 	`gorm:"type:varchar(50)"`
	Description	string	`gorm:"type:text"`
}