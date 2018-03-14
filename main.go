package main

import (
	"fmt"
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

	fmt.Printf("Start to listen on %s ...\n", c.Port)
	log.Fatal(http.ListenAndServe(":"+c.Port, router))
}
