package main

import (
	"github.com/gorilla/mux"
	"github.com/youkoulayley/api-collection/controllers"
	"github.com/youkoulayley/api-collection/middlewares"
)

var router *mux.Router

func initializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars to /cars/
	router = mux.NewRouter().StrictSlash(true)

	ApiRouter := router.PathPrefix("/api").Subrouter()

	// V1 Router
	V1Router := ApiRouter.PathPrefix("/v1").Subrouter()
	V1Router.Methods("GET").Path("/status").Name("status").HandlerFunc(controllers.HeartbeatIndex)
	V1Router.Methods("POST").Path("/login").Name("login").HandlerFunc(controllers.TokenGet)

	AdminRouter := V1Router.PathPrefix("/admin").Subrouter()
	AdminRouter.Use(middlewares.IsAuthenticate)
	AdminRouter.Use(middlewares.IsAdmin)

	RoleRouter := AdminRouter.PathPrefix("/roles").Subrouter()
	RoleRouter.Methods("GET").Path("/").Name("role.index").HandlerFunc(controllers.RoleIndex)
	RoleRouter.Methods("POST").Path("/").Name("role.create").HandlerFunc(controllers.RoleCreate)
	RoleRouter.Methods("GET").Path("/{id:[0-9]+}").Name("role.show").HandlerFunc(controllers.RoleShow)
	RoleRouter.Methods("PUT").Path("/{id:[0-9]+}").Name("role.update").HandlerFunc(controllers.RoleUpdate)
	RoleRouter.Methods("DELETE").Path("/{id:[0-9]+}").Name("role.delete").HandlerFunc(controllers.RoleDelete)

	UserRouter := V1Router.PathPrefix("/users").Subrouter()
	UserRouter.Methods("GET").Path("/").Name("user.index").HandlerFunc(controllers.UserIndex)
	UserRouter.Methods("POST").Path("/").Name("user.create").HandlerFunc(controllers.UserCreate)
	UserRouter.Methods("GET").Path("/{id:[0-9]+}").Name("user.show").HandlerFunc(controllers.UserShow)
	UserRouter.Methods("GET").Path("/{id:[0-9]+}/role").Name("user.role").HandlerFunc(controllers.UserGetRole)
	// log.Info("Routes - PUT /users/{id}")
	// router.Methods("PUT").Path("/users/{id:[0-9]+}").Name("user.update").HandlerFunc(controllers.UserUpdate)
	// log.Info("Routes - DELETE /users/{id}")
	// router.Methods("DELETE").Path("/users/{id:[0-9]+}").Name("user.delete").HandlerFunc(controllers.UserDelete)

	return router
}
