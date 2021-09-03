package httputil

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// AllowedHeaders returns a standard AllowedHeaders handlers.CORSOption.
func AllowedHeaders() handlers.CORSOption {
	return handlers.AllowedHeaders([]string{
		"Content-Type",
		"X-Requested-With",
		"Accept",
		"Accept-Language",
		"Accept-Encoding",
		"Content-Language",
		"Origin",
	})
}

// AllowedOrigins returns a handlers.CORSOption allowing all origins.
func AllowedOrigins() handlers.CORSOption {
	return handlers.AllowedOrigins([]string{"*"})
}

// AllowedMethods returns a handlers.CORSOption allowing all standard rest methods.
func AllowedMethods() handlers.CORSOption {
	return handlers.AllowedMethods([]string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"OPTIONS",
		"DELETE",
	})
}

// CORS wraps a http.Handler with the `handlers.CORSOption`s.
func CORS(h http.Handler) http.Handler {
	return handlers.CORS(AllowedOrigins(), AllowedHeaders(), AllowedMethods())(h)
}