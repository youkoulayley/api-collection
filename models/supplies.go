package models

import (
	"encoding/json"

)

// Supply struct
type Supply struct {
	Model
	SupplyTypeID int
	SupplyType   SupplyType
	Object       json.RawMessage `gorm:"type:json"`
	Videos       []Video         `gorm:"many2many:supply_video"`
}
