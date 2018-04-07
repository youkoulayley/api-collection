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
	log.Info("Routes - POST /roles")
	router.Methods("POST").Path("/roles").Name("role.create").HandlerFunc(controllers.RoleCreate)
	log.Info("Routes - GET /roles/{id}")
	router.Methods("GET").Path("/roles/{id}").Name("role.show").HandlerFunc(controllers.RoleShow)
	log.Info("Routes - PUT /roles/{id}")
	router.Methods("PUT").Path("/roles/{id}").Name("role.update").HandlerFunc(controllers.RoleUpdate)
	log.Info("Routes - DELETE /roles/{id}")
	router.Methods("DELETE").Path("/roles/{id}").Name("role.delete").HandlerFunc(controllers.RoleDelete)

	return router
}

// Router is the getter
func Router() *mux.Router {
	return router
}
