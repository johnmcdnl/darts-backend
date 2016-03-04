package utils

import (
	"encoding/json"
	"net/http"
)

func StandardContentTypeHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func DecodeRequestBodyToType(r *http.Request, obj interface{}) {
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&obj)
}
