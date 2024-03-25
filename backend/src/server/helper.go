package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func extractDecodeRequestBody[T any](r *http.Request) (T, error) {
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
