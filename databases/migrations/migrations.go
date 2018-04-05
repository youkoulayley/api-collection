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
	bootstrap.Db().AutoMigrate(&models.Role{},
							   &models.User{},
							   &models.Channel{},
							   &models.Playlist{},
							   &models.SupplyType{},
							   &models.Supply{},
							   &models.Video{})

	log.Info("Migrations - add foreign keys")
	gormfk.BelongsToFIndex(bootstrap.Db(), &models.User{}, &models.Role{})
	gormfk.BelongsToFIndex(bootstrap.Db(), &models.Channel{}, &models.User{})
	gormfk.BelongsToFIndex(bootstrap.Db(), &models.Supply{}, &models.SupplyType{})
	gormfk.BelongsToFIndex(bootstrap.Db(), &models.Playlist{}, &models.Channel{})
	gormfk.Many2ManyFIndex(bootstrap.Db(), &models.Channel{}, &models.User{})
	gormfk.Many2ManyFIndex(bootstrap.Db(), &models.Playlist{}, &models.Video{})
	gormfk.Many2ManyFIndex(bootstrap.Db(), &models.Supply{}, &models.Video{})

	log.Info("Migrations - end")
}
