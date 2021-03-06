package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ToJson(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func WriteObjectToResponse(w http.ResponseWriter, objType interface{}) {
	json, err := ToJson(objType)
	WriteJsonToResponse(w, json, err)
}

func WriteJsonToResponse(w http.ResponseWriter, json []byte, err error) {
	StandardContentTypeHeader(w)
	if err != nil {
		WriteInternalServerError(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
	log.Print(err)
	w.WriteHeader(http.StatusInternalServerError)
}
