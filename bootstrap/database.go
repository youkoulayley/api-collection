package bootstrap

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import specific
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
	"time"
)

var db *gorm.DB

// OpenDB create the connection to database
func OpenDB(c *models.Conf) {
	var err error

	d := c.Database

	connString := "postgres://" + d.User + ":" + d.Password + "@" + d.Host + ":" + d.Port + "/" + d.Database + "?sslmode=disable"

	db, err = gorm.Open("postgres", connString)
	if err != nil {
		log.Error("Database - Failed to open connection : ", err.Error())
	} else {
		log.Info("Database - Configuration is right")
	}
	err = db.DB().Ping()
	for err != nil {
		log.Error("Database - Failed to connect : ", err.Error())
		log.Debug("Sleep 30s before retry")
		time.Sleep(30 * time.Second)
		err = db.DB().Ping()
	}
	log.Info("Database - Connected")
}

// Db Getter for db var
func Db() *gorm.DB {
	return db
}
