package models

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Username string    `gorm:"type:varchar(20);unique_index;not null" json:"username,omitempty"`
	Email    string    `gorm:"type:varchar(50);unique_index;not null" json:"email,omitempty"`
	Password string    `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Role     Role      `json:"role"`
	RoleID   uint      `json:"role_id,omitempty"`
	Channels []Channel `gorm:"many2many:channel_user"`
}
