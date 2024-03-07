package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func serverHealthy() bool {
	// TODO: Include database health check
	return true
}

// HasPermission checks whether our claims have a specific permission.
// In our case, since we are using this to check if user is admin, will be checking for approve:pipeline_step permission
func (c CustomClaims) HasPermission(expectedPermission string) bool {
	result := strings.Split(c.Permissions, ",")
	for i := range result {
		if result[i] == expectedPermission {
			return true
		}
	}

	return false
}
