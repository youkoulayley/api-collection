package migrations

import (
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
)

// LaunchMigrations launch all migrations that you want
func LaunchMigrations() {
	log.Info("Migrations - start")
	bootstrap.Db().AutoMigrate(&models.Role{})
	bootstrap.Db().AutoMigrate(&models.User{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "RESTRICT")
}
