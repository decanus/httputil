package httputil

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// NotFoundHandler handles 404 responses
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	JsonError(w, http.StatusNotFound, "not found")
}

// NotAllowedHandler handles 405 responses
func NotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
	JsonError(w, http.StatusMethodNotAllowed, "not allowed")
}

// JsonError writes an Error to the ResponseWriter with the provided information.
func JsonError(w http.ResponseWriter, responseCode int, msg string) {
	type ErrorResponse struct {
		Message string    `json:"message"`
	}

	w.WriteHeader(responseCode)

	err := JsonEncode(w, ErrorResponse{Message: msg})
	if err != nil {
		log.Printf("failed to encode response: %s", err.Error())
	}
}

// JsonSuccess writes a success message to the writer.
func JsonSuccess(w http.ResponseWriter) {
	type SuccessResponse struct {
		Success bool `json:"success"`
	}

	w.WriteHeader(200)
	err := JsonEncode(w, SuccessResponse{Success: true})
	if err != nil {
		log.Printf("failed to encode response: %s", err.Error())
	}
}

// JsonEncode marshals an interface and writes it to the response.
func JsonEncode(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// GetInt returns an integer value from a URL query.
func GetInt(v url.Values, key string, defaultValue int) int {
	str := v.Get(key)
	if str == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}

	return val
}