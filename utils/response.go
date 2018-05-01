package utils

import (
	"net/http"
	"strconv"
)

// WriteJSON writes json to the response writer setting the correct headers
func WriteJSON(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

// InternalServerErrorWriter writes JSON for internal server errors
func InternalServerErrorWriter(w http.ResponseWriter) {
	WriteJSON(w, http.StatusInternalServerError,
		[]byte(`{"success": false, "message": "Something went wrong, please try again"}`))
}
