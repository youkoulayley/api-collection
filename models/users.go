package models

// User struct
type User struct {
	Model
	Username string    `gorm:"type:varchar(20);unique_index;not null" json:"username,omitempty"`
	Email    string    `gorm:"type:varchar(50);unique_index;not null" json:"email,omitempty"`
	Password string    `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	RoleID   int       `json:"role_id,omitempty"`
	Role     Role      `json:"role,omitempty"`
	Channels []Channel `gorm:"many2many:channel_user"`
}
