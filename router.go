package main

import (
	"github.com/youkoulayley/api-collection/controllers"

	"github.com/gorilla/mux"
)

// InitializeRouter do what is name tell he does
func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/paintcans").Name("Index").HandlerFunc(controllers.PaintCansIndex)
	router.Methods("POST").Path("/paintcans").Name("Create").HandlerFunc(controllers.PaintCansCreate)
	router.Methods("GET").Path("/paintcans/{id}").Name("Show").HandlerFunc(controllers.PaintCansShow)
	router.Methods("PUT").Path("/paintcans/{id}").Name("Update").HandlerFunc(controllers.PaintCansUpdate)
	router.Methods("DELETE").Path("/paintcans/{id}").Name("Delete").HandlerFunc(controllers.PaintCansDelete)
	return router
}
