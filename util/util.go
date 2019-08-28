package util

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"status": "OK", "data": data}
	json.NewEncoder(w).Encode(response)
}

func ResponseError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	response := map[string]interface{}{"status": "NG", "message": err.Error()}
	json.NewEncoder(w).Encode(response)
}
