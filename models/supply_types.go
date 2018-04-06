package models

import "github.com/jinzhu/gorm"

type SupplyType struct {
	gorm.Model
	Name        string `gorm:"type:varchar(30);unique_index"`
	Description string `gorm:"type:text"`
}
