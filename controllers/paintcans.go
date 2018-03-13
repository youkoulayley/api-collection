package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/youkoulayley/api-collection/models"

	"github.com/gorilla/mux"
)

// PaintCansIndex is the main endpoint for paint cans
func PaintCansIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.AllPaintCans())
}

// PaintCansCreate create a new paint can
func PaintCansCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var paintcan models.PaintCan

	err = json.Unmarshal(body, &paintcan)

	if err != nil {
		log.Fatal(err)
	}

	models.NewPaintCan(&paintcan)

	json.NewEncoder(w).Encode(paintcan)
}

// PaintCansShow get one pain can in the database
func PaintCansShow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	paintcan := models.FindPaintCanByID(id)

	json.NewEncoder(w).Encode(paintcan)
}

//PaintCansUpdate update a record
func PaintCansUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	paintcan := models.FindPaintCanByID(id)

	err = json.Unmarshal(body, &paintcan)
	if err != nil {
		log.Fatal(err)
	}

	models.UpdatePaintCan(paintcan)

	json.NewEncoder(w).Encode(paintcan)
}

// PaintCansDelete delete a record
func PaintCansDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	// strconv.Atoi is shorthand for ParseInt
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	err = models.DeletePaintCanByID(id)
	if err != nil {
		log.Fatal(err)
	}
}
