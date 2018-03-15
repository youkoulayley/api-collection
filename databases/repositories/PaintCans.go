package repositories

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
)

// NewPaintCan create a new paint can
func NewPaintCan(pc *models.PaintCan) *models.PaintCan {
	if pc == nil {
		log.Error(pc)
	}
	pc.CreatedAt = time.Now()
	pc.UpdatedAt = time.Now()

	query, err := bootstrap.Db().Prepare("INSERT INTO paintcans (manufacturer, color, created_at, updated_at) VALUES (?,?,?,?)")
	if err != nil {
		log.Error(err.Error())
	}

	stmt, err := query.Exec(pc.Manufacturer, pc.Color, pc.CreatedAt, pc.UpdatedAt)
	if err != nil {
		log.Error(err.Error())
	}

	lastinsertid, err := stmt.LastInsertId()
	if err != nil {
		log.Error(err.Error())
	}

	err = bootstrap.Db().QueryRow("SELECT * FROM paintcans WHERE id = ?", lastinsertid).Scan(&pc.ID, &pc.Manufacturer, &pc.Color, &pc.CreatedAt, &pc.UpdatedAt)
	if err != nil {
		log.Error(err.Error())
	}
	return pc
}

// FindPaintCanByID find a paint can in table
func FindPaintCanByID(id int) *models.PaintCan {
	var pc models.PaintCan

	row := bootstrap.Db().QueryRow("SELECT * FROM paintcans WHERE id = ?;", id)
	err := row.Scan(&pc.ID, &pc.Manufacturer, &pc.Color, &pc.CreatedAt, &pc.UpdatedAt)

	if err != nil {
		log.Debug(err.Error())
	}

	return &pc
}

// AllPaintCans get all paint cans
func AllPaintCans() *models.PaintCans {
	var pcs models.PaintCans

	rows, err := bootstrap.Db().Query("SELECT * FROM paintcans")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var pc models.PaintCan

		err := rows.Scan(&pc.ID, &pc.Manufacturer, &pc.Color, &pc.CreatedAt, &pc.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		pcs = append(pcs, pc)
	}

	return &pcs
}

// UpdatePaintCan update a paint can in table
func UpdatePaintCan(pc *models.PaintCan) {
	pc.UpdatedAt = time.Now()

	stmt, err := bootstrap.Db().Prepare("UPDATE paintcans SET manufacturer=?, color=?, updated_at=? WHERE id=?;")

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
	stmt, err := bootstrap.Db().Prepare("DELETE FROM paintcans WHERE id=?;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
