package httputil_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/decanus/httputil"
)

func TestGetInt(t *testing.T) {
	var tests = []struct {
		value        string
		expected     int
		defaultValue int
	}{
		{
			"bad",
			10,
			10,
		},
		{
			"1",
			1,
			10,
		},
		{
			"",
			10,
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {

			values := url.Values{}
			values.Set("key", tt.value)

			result := httputil.GetInt(values, "key", tt.defaultValue)
			if result != tt.expected {
				t.Fatalf("expected %d does not match actual %d", tt.expected, result)
			}
		})
	}
}

func TestJsonEncode(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		httputil.JsonSuccess(writer)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != 200 {
		t.Fatalf("returned unexpected status: %d", rr.Code)
	}

	expected := "{\"success\":true}\n"
	if rr.Body.String() != expected {
		t.Fatalf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("returned unexpected content type: %s", rr.Header().Get("Content-Type"))
	}
}

func TestNotFoundHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		httputil.NotFoundHandler(writer, request)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("returned unexpected status: %d", rr.Code)
	}

	expected := "{\"message\":\"not found\"}\n"
	if rr.Body.String() != expected {
		t.Fatalf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("returned unexpected content type: %s", rr.Header().Get("Content-Type"))
	}
}

func TestNotAllowedHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		httputil.NotAllowedHandler(writer, request)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("returned unexpected status: %d", rr.Code)
	}

	expected := "{\"message\":\"not allowed\"}\n"
	if rr.Body.String() != expected {
		t.Fatalf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("returned unexpected content type: %s", rr.Header().Get("Content-Type"))
	}
}
