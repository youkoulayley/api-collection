package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import Mysql specific
)

// CreatePaintCansTable create the table if not exists
func CreatePaintCansTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS paintcans(id serial, manufacturer varchar(50), color varchar(50), created_at timestamp NULL DEFAULT NULL, updated_at timestamp NULL DEFAULT NULL, constraint pk primary key(id))")

	if err != nil {
		fmt.Printf("Failed to create table : %s", err.Error())
	}
}
