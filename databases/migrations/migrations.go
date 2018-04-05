package migrations

import (
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/gormfk"
)

// LaunchMigrations launch all migrations that you want
func LaunchMigrations() {
	log.Info("Migrations - start")
	bootstrap.Db().AutoMigrate(&models.Role{})
	bootstrap.Db().AutoMigrate(&models.User{})
	bootstrap.Db().AutoMigrate(&models.Channel{})

	log.Info("Migrations - add foreign keys")
	gormfk.BelongsToFIndex(bootstrap.Db(), &models.User{}, &models.Role{})
	gormfk.BelongsToFIndex(bootstrap.Db(), &models.Channel{}, &models.User{})
	gormfk.Many2ManyFIndex(bootstrap.Db(), &models.Channel{}, &models.User{})

	log.Info("Migrations - end")
}
