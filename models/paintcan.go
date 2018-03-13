package models

import (
	"log"
	"time"

	"github.com/youkoulayley/api-collection/databases"
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

	err := databases.Db().QueryRow("INSERT INTO paintcans (manufacturer, color, created_at, updated_at) VALUES (?,?,?,?)", pc.Manufacturer, pc.Color, pc.CreatedAt, pc.UpdatedAt).Scan(&pc.ID)

	if err != nil {
		log.Fatal(err)
	}
}

// FindPaintCanByID find a paint can in table
func FindPaintCanByID(id int) *PaintCan {
	var cp PaintCan

	row := databases.Db().QueryRow("SELECT * FROM paintcans WHERE id = ?;", id)
	err := row.Scan(&cp.ID, &cp.Manufacturer, &cp.Color, &cp.CreatedAt, &cp.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &cp
}

// AllPaintCans get all paint cans
func AllPaintCans() *PaintCans {
	var pcs PaintCans

	rows, err := databases.Db().Query("SELECT * FROM paintcans")

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

	stmt, err := databases.Db().Prepare("UPDATE paintcans SET manufacturer=?, color=?, updated_at=? WHERE id=?;")

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
	stmt, err := databases.Db().Prepare("DELETE FROM paintcans WHERE id=?;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
