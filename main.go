package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/databases/migrations"
)

func main() {
	c := bootstrap.GetConf()
	bootstrap.InitLogs(c)
	bootstrap.OpenDB(c)
	migrations.LaunchMigrations()

	router := InitializeRouter()

	log.Info("Start to listen on ", c.Port, " ...")
	log.Fatal(http.ListenAndServe(":"+c.Port, router))
}
