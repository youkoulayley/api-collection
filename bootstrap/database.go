package bootstrap

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import Mysql specific
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
)

var db *sql.DB

// OpenDB create the connection to database
func OpenDB(c *models.Conf) {
	var err error

	d := c.Database

	db, err = sql.Open("mysql", d.MysqlUser+":"+d.MysqlPassword+"@tcp("+d.MysqlHost+":"+d.MysqlPort+")/"+d.MysqlDatabase+"?charset=utf8&parseTime=True")
	if err != nil {
		log.Error("Database - Failed to open connection : ", err.Error())
	} else {
		log.Info("Database - Successfully connected")
	}

	err = db.Ping()
	if err != nil {
		log.Error("Database - Failed to connect : ", err.Error())
	} else {
		log.Info("Database - Joinable")
	}
}

// Db Getter for db var
func Db() *sql.DB {
	return db
}
