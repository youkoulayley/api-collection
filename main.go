package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/youkoulayley/api-collection/bootstrap"
)

func main() {
	c := bootstrap.GetConf()
	bootstrap.InitLogs()
	bootstrap.OpenDB(c)

	router := InitializeRouter()

	fmt.Printf("Start to listen on %s ...\n", c.Port)
	log.Fatal(http.ListenAndServe(":"+c.Port, router))
}
