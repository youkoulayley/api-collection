package models

import (
	"github.com/jinzhu/gorm"
)

// Channel struct
type Channel struct {
	gorm.Model
	Name        string `json:"name,omitempty" gorm:"type:varchar(20);unique_index"`
	Description string `json:"description,omitempty" gorm:"type:text"`
	URL         string `json:"url,omitempty" gorm:"type:varchar(255);unique_index"`
	UserID      int    `json:"user_id,omitempty"`
	User        User   `json:"user,omitempty"`
	Suscribers  []User `gorm:"many2many:channel_user"`
}
