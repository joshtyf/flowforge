package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

// Decode the request body into a struct of type T. Request body buffer will be refilled after decoding.
func decode[T any](r *http.Request) (T, error) {
	var v T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	if err := json.NewDecoder(io.NopCloser(bytes.NewBuffer(body))).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))
	return v, nil
}

func serverHealthy() bool {
	// TODO: Include database health check
	return true
}

// Extracts a query parameter value from a URL query string based on the key.
//
// emptyStringAllowed determines whether an empty string is expected as a valid value.
//
// If the key is not found, and emptyStringAllowed is false, returns the defaultValue with no error.
//
// If the converter function returns an error, returns the defaultValue with the error.
// This is the only case extractQueryParam will return an error.
//
// Else, the value is returned after it has been converted using the converter function with no error.
func extractQueryParam[T any](queryParams url.Values, key string, emptyStringAllowed bool, defaultValue T, converter func(string) (T, error)) (T, error) {
	value := queryParams.Get(key)
	if value == "" && !emptyStringAllowed {
		return defaultValue, nil
	}
	ret, err := converter(value)
	if err != nil {
		return defaultValue, err
	}
	return ret, nil
}

// Helper function to convert a string to an integer. Wrapper around strconv.Atoi.
func integerConverter(s string) (int, error) {
	return strconv.Atoi(s)
}
