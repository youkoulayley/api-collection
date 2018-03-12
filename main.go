package main

import (
	"collection/config"
	"collection/models"
	"log"
	"net/http"
)

func main() {
	config.OpenDB()
	router := InitializeRouter()

	// Populate database
	models.NewPaintCan(&models.PaintCan{Manufacturer: "Citadel", Color: "White Scar"})

	log.Fatal(http.ListenAndServe(":8080", router))
}
