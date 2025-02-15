package utils

import (
	"encoding/json"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578
	http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	type response struct {
		Data any `json:"data"`
	}
	return WriteJSON(w, status, &response{Data: data})
}

func WriteJSONErrorResponse(w http.ResponseWriter, status int, message string) error {
	type response struct {
		Error any `json:"error"`
	}
	return WriteJSON(w, status, &response{Error: message})
}
