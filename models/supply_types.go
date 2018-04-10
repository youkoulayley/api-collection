package models

// SupplyType struct
type SupplyType struct {
	Model
	Name        string `gorm:"type:varchar(30);unique_index"`
	Description string `gorm:"type:text"`
}
