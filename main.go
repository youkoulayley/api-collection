package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/youkoulayley/api-collection/config"
	"github.com/youkoulayley/api-collection/databases"
)

func main() {
	conf := config.GetConf()
	databases.OpenDB(conf.Database.MysqlUser, conf.Database.MysqlPassword, conf.Database.MysqlHost, conf.Database.MysqlDatabase, conf.Database.MysqlPort)

	router := InitializeRouter()

	fmt.Printf("Start to listen on %s ...\n", conf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.Port, router))
}
