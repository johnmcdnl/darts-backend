package utils

import (
	"encoding/json"
	"net/http"
	"log"
)

func ToJson(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func WriteObjectToResponse(w http.ResponseWriter, myVars interface{}) {
	json, err := ToJson(myVars)
	WriteJsonToResponse(w, json, err)
}

func WriteJsonToResponse(w http.ResponseWriter, json []byte, err error) {
	StandardContentTypeHeader(w)
	if (err != nil) {
		WriteInternalServerError(w, err)
	}else {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
	log.Print(err)
	w.WriteHeader(http.StatusInternalServerError)
}