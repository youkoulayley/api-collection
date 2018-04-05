package models

import "github.com/jinzhu/gorm"

// User struct
type User struct {
	gorm.Model
	Username string    `gorm:"type:varchar(20);unique_index" json:"username,omitempty"`
	Email    string    `gorm:"type:varchar(50);unique_index" json:"email,omitempty"`
	Password string    `gorm:"type:varchar(255)" json:"password,omitempty"`
	RoleID   int       `json:"role_id,omitempty"`
	Role     Role      `json:"role,omitempty"`
	Channels []Channel `gorm:"many2many:channel_user"`
}
