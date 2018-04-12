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

	log.Info("Routes - POST /token")
	router.Methods("POST").Path("/token").Name("token").HandlerFunc(controllers.TokenGet)

	// ROLES
	log.Info("Routes - GET /roles")
	router.Methods("GET").Path("/roles").Name("role.index").HandlerFunc(controllers.RoleIndex)
	log.Info("Routes - POST /roles")
	router.Methods("POST").Path("/roles").Name("role.create").HandlerFunc(controllers.RoleCreate)
	log.Info("Routes - GET /roles/{id}")
	router.Methods("GET").Path("/roles/{id:[0-9]+}").Name("role.show").HandlerFunc(controllers.RoleShow)
	log.Info("Routes - PUT /roles/{id}")
	router.Methods("PUT").Path("/roles/{id:[0-9]+}").Name("role.update").HandlerFunc(controllers.RoleUpdate)
	log.Info("Routes - DELETE /roles/{id}")
	router.Methods("DELETE").Path("/roles/{id:[0-9]+}").Name("role.delete").HandlerFunc(controllers.RoleDelete)

	// USERS
	log.Info("Routes - GET /users")
	router.Methods("GET").Path("/users").Name("user.index").HandlerFunc(controllers.UserIndex)
	log.Info("Routes - POST /users")
	router.Methods("POST").Path("/users").Name("user.create").HandlerFunc(controllers.UserCreate)
	log.Info("Routes - GET /users/{id}")
	router.Methods("GET").Path("/users/{id:[0-9]+}").Name("user.show").HandlerFunc(controllers.UserShow)
	log.Info("Routes - Get /users/{id}/role/")
	router.Methods("GET").Path("/users/{id:[0-9]+}/role").Name("user.role").HandlerFunc(controllers.UserGetRole)
	// log.Info("Routes - PUT /users/{id}")
	// router.Methods("PUT").Path("/users/{id:[0-9]+}").Name("user.update").HandlerFunc(controllers.UserUpdate)
	// log.Info("Routes - DELETE /users/{id}")
	// router.Methods("DELETE").Path("/users/{id:[0-9]+}").Name("user.delete").HandlerFunc(controllers.UserDelete)

	return router
}

// Router is the getter
func Router() *mux.Router {
	return router
}
