package utils

import (
	"encoding/json"
	"fmt"
	"football-simulation/types"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)

}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func WriteSuccess(w http.ResponseWriter, status int, data interface{}) {
	response := types.Response{
		Status: "success",
		Data:   data,
	}
	WriteJSON(w, status, response)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	response := types.Response{
		Status:  "error",
		Error:   err,
		Message: err.Error(),
	}
	WriteJSON(w, status, response)
}
