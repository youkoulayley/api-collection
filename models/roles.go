package models

import "github.com/jinzhu/gorm"

// Role struct
type Role struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);unique_index;not null" json:"name,omitempty"`
	Description string `gorm:"type:text" json:"description,omitempty"`
}
