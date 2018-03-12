package config

import (
	"collection/databases"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import Mysql specific
)

var db *sql.DB

// OpenDB create the connection to database
func OpenDB() {
	var err error

	db, err = sql.Open("mysql", "root:titoun@tcp(192.168.1.43:3306)/collection")
	if err != nil {
		fmt.Printf("Failed to open connection to database : %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to database : %s", err.Error())
	}

	// Create painting table if not exist
	databases.CreatePaintCansTable(db)
}

// Db Getter for db var
func Db() *sql.DB {
	return db
}
