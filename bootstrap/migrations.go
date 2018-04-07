package bootstrap

import (
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/gormfk"
)

// LaunchMigrations launch all migrations that you want
func LaunchMigrations() {
	log.Info("Migrations - start")
	Db().AutoMigrate(&models.Role{},
					 &models.User{},
					 &models.Channel{},
					 &models.Playlist{},
					 &models.SupplyType{},
				     &models.Supply{},
				     &models.Video{})

	log.Info("Migrations - add foreign keys")
	gormfk.BelongsTo(Db(), &models.User{}, &models.Role{})
	gormfk.BelongsTo(Db(), &models.Channel{}, &models.User{})
	gormfk.BelongsTo(Db(), &models.Supply{}, &models.SupplyType{})
	gormfk.BelongsTo(Db(), &models.Playlist{}, &models.Channel{})
	gormfk.Many2Many(Db(), &models.Channel{}, &models.User{})
	gormfk.Many2Many(Db(), &models.Playlist{}, &models.Video{})
	gormfk.Many2Many(Db(), &models.Supply{}, &models.Video{})

	log.Info("Migrations - end")
}
