package targets

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/johnmcdnl/darts-backend/utils"
	"strconv"
)

func TargetRouter(r *mux.Router) {
	r.HandleFunc("/targets/create", createMatchOptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/targets/create", createMatchHandler).Methods("POST")

	r.HandleFunc("/targets", allMatchesHandler).Methods("GET")
	r.HandleFunc("/targets/user/{userId}", allMatchesByUserHandler).Methods("GET")

	r.HandleFunc("/targets/targetname/{targetName}", allMatchesByTargetNameHandler).Methods("GET")
	r.HandleFunc("/targets/targetname/{targetName}/user/{userId}", allMatchesByUserAndTargetNameHandler).Methods("GET")
	r.HandleFunc("/targets/targetname/{targetName}/user/{userId}/average", allMatchesByUserAndTargetNameAverageHandler).Methods("GET")

	r.HandleFunc("/targets/areatype/{areaType}", allMatchesByAreaTypeHandler).Methods("GET")
	r.HandleFunc("/targets/areatype/{areaType}/user/{userId}", allMatchesByUserAndAreaTypeHandler).Methods("GET")
}

func createMatchOptionsHandler(w http.ResponseWriter, r *http.Request) {
	//TODO something better
	w.WriteHeader(http.StatusOK)
}

func createMatchHandler(w http.ResponseWriter, r *http.Request) {

}

func allMatchesHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteObjectToResponse(w, allMatches())
}

func allMatchesByUserHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["userId"])
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
	}else {
		utils.WriteObjectToResponse(w, allMatchesForUser(userId))
	}

}

func allMatchesByTargetNameHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteObjectToResponse(w, allMatchesByTargetName(mux.Vars(r)["targetName"]))
}

func allMatchesByUserAndTargetNameHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["userId"])
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
	}
	utils.WriteObjectToResponse(w, allMatchesByUserAndTarget(userId, mux.Vars(r)["targetName"]))
}

func allMatchesByUserAndTargetNameAverageHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["userId"])
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
	}
	utils.WriteObjectToResponse(w, allMatchesByUserAndTargetAverage(userId, mux.Vars(r)["targetName"]))
}

func allMatchesByAreaTypeHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteObjectToResponse(w, allMatchesByAreaType(mux.Vars(r)["areaType"]))
}

func allMatchesByUserAndAreaTypeHandler(w http.ResponseWriter, r *http.Request) {

}
