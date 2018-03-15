package main

import (
	"github.com/youkoulayley/api-collection/controllers"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var router *mux.Router

// InitializeRouter do what is name tell he does
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	log.Info("Routes - GET /status")
	router.Methods("GET").Path("/status").Name("Status").HandlerFunc(controllers.HeartbeatIndex)

	log.Info("Routes - /paintcans GET")
	router.Methods("GET").Path("/paintcans").Name("Index").HandlerFunc(controllers.PaintCansIndex)
	log.Info("Routes - /paintcans POST")
	router.Methods("POST").Path("/paintcans").Name("Create").HandlerFunc(controllers.PaintCansCreate)
	log.Info("Routes - /paintcans/{id} GET")
	router.Methods("GET").Path("/paintcans/{id}").Name("Show").HandlerFunc(controllers.PaintCansShow)
	log.Info("Routes - /paintcans/{id} PUT")
	router.Methods("PUT").Path("/paintcans/{id}").Name("Update").HandlerFunc(controllers.PaintCansUpdate)
	log.Info("Routes - /paintcans/{id} DELETE")
	router.Methods("DELETE").Path("/paintcans/{id}").Name("Delete").HandlerFunc(controllers.PaintCansDelete)

	return router
}

func Router() *mux.Router {
	return router
}
