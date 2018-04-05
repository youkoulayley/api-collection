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
	router = mux.NewRouter().StrictSlash(true)

	log.Info("Routes - GET /status")
	router.Methods("GET").Path("/status").Name("status").HandlerFunc(controllers.HeartbeatIndex)

	log.Info("Routes - GET /roles")
	router.Methods("GET").Path("/roles").Name("role.index").HandlerFunc(controllers.RoleIndex)
	// log.Info("Routes - POST /paintcans")
	// router.Methods("POST").Path("/paintcans").Name("Create").HandlerFunc(controllers.PaintCansCreate)
	// log.Info("Routes - GET /paintcans/{id}")
	// router.Methods("GET").Path("/paintcans/{id}").Name("Show").HandlerFunc(controllers.PaintCansShow)
	// log.Info("Routes - PUT /paintcans/{id}")
	// router.Methods("PUT").Path("/paintcans/{id}").Name("Update").HandlerFunc(controllers.PaintCansUpdate)
	// log.Info("Routes - DELETE /paintcans/{id}")
	// router.Methods("DELETE").Path("/paintcans/{id}").Name("Delete").HandlerFunc(controllers.PaintCansDelete)

	return router
}

// Router is the getter
func Router() *mux.Router {
	return router
}
