package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data any, statusCode int) {

	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, msg any, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
