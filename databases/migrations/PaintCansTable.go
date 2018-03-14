package migrations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import Mysql specific
	log "github.com/sirupsen/logrus"
)

// CreatePaintCansTable create the table if not exists
func CreatePaintCansTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS paintcans(id serial, manufacturer varchar(50), color varchar(50), created_at timestamp NULL DEFAULT NULL, updated_at timestamp NULL DEFAULT NULL, constraint pk primary key(id))")

	if err != nil {
		log.Error("Failed to create table : ", err.Error())
	} else {
		log.Info("Migrations - PaintCans successfuly migrated")
	}
}
