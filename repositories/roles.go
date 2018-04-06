package repositories

import (
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
)

// RolesGetAll fetch all roles
func RolesGetAll() []models.Role {
	var r []models.Role

	bootstrap.Db().Find(&r)

	return r
}

// RoleCreate create a new role
func RoleCreate(r *models.Role) {
	if r == nil {
		log.Error(r)
	}

	// Check if role already exists
	var count uint
	bootstrap.Db().Where("name = ?", r.Name).Find(&models.Role{}).Count(&count)

	if count == 0 {
		err := bootstrap.Db().Create(&r).Error
		if err != nil {
			log.Error(err)
		}
	}
}

//
//// FindPaintCanByID find a paint can in table
//func FindPaintCanByID(id int) *models.PaintCan {
//	var pc models.PaintCan
//
//	row := bootstrap.Db().QueryRow("SELECT * FROM paintcans WHERE id = ?;", id)
//	err := row.Scan(&pc.ID, &pc.Manufacturer, &pc.Color, &pc.CreatedAt, &pc.UpdatedAt)
//
//	if err != nil {
//		log.Debug(err.Error())
//	}
//
//	return &pc
//}
//
//
//	return &pcs
//}
//
//// UpdatePaintCan update a paint can in table
//func UpdatePaintCan(pc *models.PaintCan) {
//	pc.UpdatedAt = time.Now()
//
//	stmt, err := bootstrap.Db().Prepare("UPDATE paintcans SET manufacturer=?, color=?, updated_at=? WHERE id=?;")
//
//	if err != nil {
//		log.Debug(err)
//	}
//
//	_, err = stmt.Exec(pc.Manufacturer, pc.Color, pc.UpdatedAt, pc.ID)
//
//	if err != nil {
//		log.Debug(err)
//	}
//}
//
//// DeletePaintCanByID delete one paintcan
//func DeletePaintCanByID(id int) error {
//	stmt, err := bootstrap.Db().Prepare("DELETE FROM paintcans WHERE id=?;")
//
//	if err != nil {
//		log.Debug(err)
//	}
//
//	_, err = stmt.Exec(id)
//
//	return err
//}
