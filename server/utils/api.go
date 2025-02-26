package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError struct {
	Error   error
	Message string
	Status  int
}

const ContentTypeJSON = "application/json"

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, e *ApiError) {
	WriteJSON(w, e.Status, map[string]any{"error": e.Message, "status": e.Status})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return errors.New("Missing request body")
	}

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(v)
}
