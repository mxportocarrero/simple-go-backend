package utils

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type GenericResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func WriteError(w http.ResponseWriter, code int, message string, data interface{}) {
	response := GenericResponse{
		Status: "error",
		Error:  message,
		Data:   data,
	}

	WriteJSON(w, code, response)
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	response := GenericResponse{
		Status: "success",
		Data:   data,
	}

	WriteJSON(w, code, response)
}

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.WithError(err).Warn("Error writing response")
	}
}
