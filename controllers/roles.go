package controllers

import (
	"net/http"
	"github.com/youkoulayley/api-collection/repositories"
	log "github.com/sirupsen/logrus"
)

func RoleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	log.Info(repositories.RolesGetAll())
}

//func PaintCansIndex(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json;charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//
//	json.NewEncoder(w).Encode(repositories.AllPaintCans())
//}
//
//// PaintCansCreate create a new paint can
//func PaintCansCreate(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json;charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//
//	body, err := ioutil.ReadAll(r.Body)
//
//	if err != nil {
//		log.Error(err)
//	}
//
//	var paintcan models.PaintCan
//
//	err = json.Unmarshal(body, &paintcan)
//	if err != nil {
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(422) // unprocessable entity
//		if err = json.NewEncoder(w).Encode(err); err != nil {
//			log.Error(err.Error())
//		}
//	} else {
//		repositories.NewPaintCan(&paintcan)
//		json.NewEncoder(w).Encode(paintcan)
//	}
//}
//
//// PaintCansShow get one pain can in the database
//func PaintCansShow(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json;charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//
//	if err != nil {
//		log.Error(err.Error())
//	}
//
//	paintcan := repositories.FindPaintCanByID(id)
//
//	if paintcan.ID == 0 {
//		json.NewEncoder(w).Encode(models.Heartbeat{Status: "Not Found", Code: 404})
//	} else {
//		json.NewEncoder(w).Encode(paintcan)
//	}
//}
//
////PaintCansUpdate update a record
//func PaintCansUpdate(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json;charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//
//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//
//	if err != nil {
//		log.Error(err.Error())
//	}
//
//	body, err := ioutil.ReadAll(r.Body)
//
//	if err != nil {
//		log.Error(err.Error())
//	}
//
//	paintcan := repositories.FindPaintCanByID(id)
//
//	err = json.Unmarshal(body, &paintcan)
//	if err != nil {
//		log.Error(err.Error())
//	}
//
//	repositories.UpdatePaintCan(paintcan)
//
//	json.NewEncoder(w).Encode(paintcan)
//}
//
//// PaintCansDelete delete a record
//func PaintCansDelete(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "application/json;charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//
//	vars := mux.Vars(r)
//
//	// strconv.Atoi is shorthand for ParseInt
//	id, err := strconv.Atoi(vars["id"])
//
//	if err != nil {
//		log.Error(err)
//	}
//
//	err = repositories.DeletePaintCanByID(id)
//	if err != nil {
//		log.Error(err.Error())
//	} else {
//		json.NewEncoder(w).Encode(models.Heartbeat{Status: "OK", Code: 200})
//	}
//}