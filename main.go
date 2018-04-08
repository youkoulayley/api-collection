package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	"github.com/youkoulayley/api-collection/bootstrap"
)

func main() {
	confPath := bootstrap.InitFlags()             // Init and parse the flags
	c := bootstrap.GetConf(confPath, "conf.json") // Get conf and store it in var

	log.Info("Start the app in v", c.Version)

	bootstrap.InitLogs(c)        	 // Init the logs
	bootstrap.OpenDB(c)           	// Init the DB
	bootstrap.LaunchMigrations()	// Launch database migration

	router := InitializeRouter() // Init router

	// CORS HEADERS
	corsAllowedOrigins := handlers.AllowedOrigins(c.Cors.AuthorizedHosts) // Authorized hosts
	corsAllowedMethods := handlers.AllowedMethods(c.Cors.AuthorizedMethods) // Authorized methods
	corsAllowedHeaders := handlers.AllowedHeaders(c.Cors.AuthorizedHeaders) // Authorized headers

	log.Info("Start to listen on ", c.Port, " ...")
	log.Fatal(http.ListenAndServe(":"+c.Port, handlers.CORS(corsAllowedOrigins, corsAllowedMethods, corsAllowedHeaders)(router))) // Serve the HTTP Server. Exit if fail.
}
