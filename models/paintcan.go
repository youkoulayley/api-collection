package models

import (
	"collection/config"
	"log"
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

// NewPaintCan create a new paint can
func NewPaintCan(pc *PaintCan) {
	if pc == nil {
		log.Fatal(pc)
	}
	pc.CreatedAt = time.Now()
	pc.UpdatedAt = time.Now()

	err := config.Db().QueryRow("INSERT INTO paintcans (manufacturer, color, created_at, updated_at) VALUES (?,$2,$3,$4)", pc.Manufacturer, pc.Color, pc.CreatedAt, pc.UpdatedAt).Scan(&pc.ID)

	if err != nil {
		log.Fatal(err)
	}
}

// FindPaintCanByID find a paint can in table
func FindPaintCanByID(id int) *PaintCan {
	var cp PaintCan

	row := config.Db().QueryRow("SELECT * FROM paintcans WHERE id = $1;", id)
	err := row.Scan(&cp.ID, &cp.Manufacturer, &cp.Color, &cp.CreatedAt, &cp.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &cp
}

// AllPaintCans get all paint cans
func AllPaintCans() *PaintCans {
	var pcs PaintCans

	rows, err := config.Db().Query("SELECT * FROM paincans")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var pc PaintCan

		err := rows.Scan(&pc.ID, &pc.Manufacturer, &pc.Color, &pc.CreatedAt, &pc.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		pcs = append(pcs, pc)
	}

	return &pcs
}

// UpdatePaintCan update a paint can in table
func UpdatePaintCan(pc *PaintCan) {
	pc.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE paintcans SET manufacturer=$1, color=$2, updated_at=$3 WHERE id=$4;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(pc.Manufacturer, pc.Color, pc.UpdatedAt, pc.ID)

	if err != nil {
		log.Fatal(err)
	}
}

// DeletePaintCanByID delete one paintcan
func DeletePaintCanByID(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM paintcans WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
