package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Supply struct {
	gorm.Model
	SupplyTypeID int
	SupplyType   SupplyType
	Object       json.RawMessage `gorm:"type:json"`
	Videos       []Video         `gorm:"many2many:supply_video"`
}
