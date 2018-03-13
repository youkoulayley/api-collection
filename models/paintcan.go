package models

import (
	"time"
)

// PaintCan contain the struct for paintcan
type PaintCan struct {
	ID           int       `json:"id"`
	Manufacturer string    `json:"manufacturer"`
	Color        string    `json:"color"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PaintCans type
type PaintCans []PaintCan
