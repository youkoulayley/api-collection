package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/databases/migrations"
)

func main() {
	confPath := flag.String("config", "./", "Path for the config file")
	flag.Parse()

	c := bootstrap.GetConf(*confPath, "conf.json")
	bootstrap.InitLogs(c)
	bootstrap.OpenDB(c)
	migrations.LaunchMigrations()

	router := InitializeRouter()

	corsObj := handlers.AllowedOrigins(c.AuthorizedHosts)

	log.Info("Start to listen on ", c.Port, " ...")
	log.Fatal(http.ListenAndServe(":"+c.Port, handlers.CORS(corsObj)(router)))
}
