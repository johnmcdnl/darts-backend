package board

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/johnmcdnl/darts-backendv2/utils"
)

func BoardRouter(r *mux.Router) {
	r.HandleFunc("/areas", allAreasHandler).Methods("GET")
	r.HandleFunc("/areas/{areaName}", areaByNameHandler).Methods("GET")
	r.HandleFunc("/areatypes", allAreaTypesHandler).Methods("GET")
	r.HandleFunc("/areatypes/{areaTypeName}", areasByAreaTypeHandler).Methods("GET")
}

func allAreasHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteObjectToResponse(w, getAllAreas())
}

func areaByNameHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteObjectToResponse(w, getAreaByName(mux.Vars(r)["areaName"]))
}

func allAreaTypesHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteObjectToResponse(w, getAllAreaTypes())
}

func areasByAreaTypeHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteObjectToResponse(w, getAreaTypeByName(mux.Vars(r)["areaTypeName"]))
}
