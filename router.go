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
	V1Router.Methods("GET").Path("/status").HandlerFunc(controllers.HeartbeatIndex).Name("status")
	V1Router.Methods("POST").Path("/login").HandlerFunc(controllers.TokenGet).Name("login")
	V1Router.Methods("POST").Path("/logout").HandlerFunc(controllers.TokenRemove).Name("logout")

	AdminRouter := V1Router.PathPrefix("/admin").Subrouter()
	AdminRouter.Use(middlewares.IsAuthenticate)
	AdminRouter.Use(middlewares.IsAdmin)

	AdminRoleRouter := AdminRouter.PathPrefix("/roles").Subrouter()
	AdminRoleRouter.Methods("GET").Path("/").Name("role.index").HandlerFunc(controllers.RoleIndex)
	AdminRoleRouter.Methods("POST").Path("/").Name("role.create").HandlerFunc(controllers.RoleCreate)
	AdminRoleRouter.Methods("GET").Path("/{id:[0-9]+}").Name("role.show").HandlerFunc(controllers.RoleShow)
	AdminRoleRouter.Methods("PUT").Path("/{id:[0-9]+}").Name("role.update").HandlerFunc(controllers.RoleUpdate)
	AdminRoleRouter.Methods("DELETE").Path("/{id:[0-9]+}").Name("role.delete").HandlerFunc(controllers.RoleDelete)

	AdminUserRouter := V1Router.PathPrefix("/users").Subrouter()
	AdminUserRouter.Methods("GET").Path("/").Name("user.index").HandlerFunc(controllers.UserIndex)
	AdminUserRouter.Methods("POST").Path("/").Name("user.create").HandlerFunc(controllers.UserCreate)
	AdminUserRouter.Methods("GET").Path("/{id:[0-9]+}").Name("user.show").HandlerFunc(controllers.UserShow)
	AdminUserRouter.Methods("GET").Path("/{id:[0-9]+}/role").Name("user.role").HandlerFunc(controllers.UserGetRole)
	// log.Info("Routes - PUT /users/{id}")
	// router.Methods("PUT").Path("/users/{id:[0-9]+}").Name("user.update").HandlerFunc(controllers.UserUpdate)
	// log.Info("Routes - DELETE /users/{id}")
	// router.Methods("DELETE").Path("/users/{id:[0-9]+}").Name("user.delete").HandlerFunc(controllers.UserDelete)

	return router
}
