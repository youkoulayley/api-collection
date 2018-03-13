package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import Mysql specific
	"github.com/youkoulayley/api-collection/databases/migrations"
	"github.com/youkoulayley/api-collection/models"
)

var db *sql.DB

// OpenDB create the connection to database
func OpenDB(c *models.Conf) {
	var err error

	d := c.Database

	db, err = sql.Open("mysql", d.MysqlUser+":"+d.MysqlPassword+"@tcp("+d.MysqlHost+":"+d.MysqlPort+")/"+d.MysqlDatabase+"?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Printf("Failed to open connection to database : %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Failed to connect to database : %s", err.Error())
	}

	// Create painting table if not exist
	migrations.CreatePaintCansTable(db)
}

// Db Getter for db var
func Db() *sql.DB {
	return db
}
