package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/controllers"
)

func main() {
	confPath := bootstrap.InitFlags()             // Init and parse the flags
	c := bootstrap.GetConf(confPath, "conf.json") // Get conf and store it in var

	controllers.JwtSalt = []byte(c.JWTSalt)

	log.Info("Start the app in v", c.Version)

	bootstrap.InitLogs(c)        // Init the logs
	bootstrap.OpenDB(c)          // Init the DB
	bootstrap.LaunchMigrations() // Launch database migration

	r := initializeRouter() // Init router

	// HEADERS
	AllowedOrigins := handlers.AllowedOrigins(c.Cors.AuthorizedHosts)   // Authorized hosts
	AllowedMethods := handlers.AllowedMethods(c.Cors.AuthorizedMethods) // Authorized methods
	AllowedHeaders := handlers.AllowedHeaders(c.Cors.AuthorizedHeaders) // Authorized headers

	log.Info("Start to listen on ", c.Port, " ...")
	log.Fatal(http.ListenAndServe(":"+c.Port, handlers.CORS(AllowedOrigins, AllowedMethods, AllowedHeaders)(r))) // Serve the HTTP Server. Exit if fail.
}
