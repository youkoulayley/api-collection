package main

import (
	"collection/controllers"

	"github.com/gorilla/mux"
)

// InitializeRouter do what is name tell he does
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/cars").Name("Index").HandlerFunc(controllers.PaintCansIndex)
	router.Methods("POST").Path("/cars").Name("Create").HandlerFunc(controllers.PaintCanCreate)
	router.Methods("GET").Path("/cars/{id}").Name("Show").HandlerFunc(controllers.PaintCansShow)
	router.Methods("PUT").Path("/cars/{id}").Name("Update").HandlerFunc(controllers.PaintCansUpdate)
	router.Methods("DELETE").Path("/cars/{id}").Name("DELETE").HandlerFunc(controllers.PaintCansDelete)
	return router
}
