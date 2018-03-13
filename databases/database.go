package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import Mysql specific
)

var db *sql.DB

// OpenDB create the connection to database
func OpenDB(u string, p string, h string, d string, P string) {
	var err error

	db, err = sql.Open("mysql", u+":"+p+"@tcp("+h+":"+P+")/"+d)
	if err != nil {
		fmt.Printf("Failed to open connection to database : %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to database : %s", err.Error())
	}

	// Create painting table if not exist
	CreatePaintCansTable(db)
}

// Db Getter for db var
func Db() *sql.DB {
	return db
}
