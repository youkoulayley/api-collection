package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/databases/migrations"
)

func main() {
	confPath := bootstrap.InitFlags()             // Init and parse the flags
	c := bootstrap.GetConf(confPath, "conf.json") // Get conf and store it in var
	bootstrap.InitLogs(c)                         // Init the logs
	bootstrap.OpenDB(c)                           // Init the DB
	migrations.LaunchMigrations()                 // Launch database migration

	router := InitializeRouter() // Init router

	corsObj := handlers.AllowedOrigins(c.AuthorizedHosts) // Init CORS Header

	log.Info("Start to listen on ", c.Port, " ...")
	log.Fatal(http.ListenAndServe(":"+c.Port, handlers.CORS(corsObj)(router))) // Serve the HTTP Server. Exit if fail.
}
