package migrations

import (
	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/bootstrap"
)

// LaunchMigrations launch all migrations that you want
func LaunchMigrations() {
	log.Info("Migrations - start")
	CreatePaintCansTable(bootstrap.Db())
}
