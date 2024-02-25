package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/joshtyf/flowforge/src/api"
	"github.com/joshtyf/flowforge/src/authenticator"
	"github.com/joshtyf/flowforge/src/logger"
	"golang.org/x/oauth2"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		accessToken := r.Header.Get("Authorization")

		var profile map[string]interface{}

		authToken := &oauth2.Token{AccessToken: accessToken}
		if accessToken == "" {
			authError := &api.HandlerError{Message: "authorization token missing", Code: http.StatusUnauthorized}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		auth, err := authenticator.New()
		if err != nil {
			logger.Error("[Authorization] Failed to initialize authenticator", map[string]interface{}{"err": err})
			authError := &api.HandlerError{Message: api.ErrInternalServerError.Error(), Code: http.StatusInternalServerError}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		idToken, err := auth.VerifyIDToken(r.Context(), authToken)
		if err != nil {
			logger.Error("[Authorization] Unable to verify token", map[string]interface{}{"err": err})
			authError := &api.HandlerError{Message: "unable to verify token", Code: http.StatusUnauthorized}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		err = idToken.Claims(&profile)
		if err != nil {
			logger.Error("[Authorization] Unable retrieve profile", map[string]interface{}{"err": err})
			authError := &api.HandlerError{Message: "unable to retrieve profile", Code: http.StatusInternalServerError}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		// Call the next middleware function or final handler
		next.ServeHTTP(w, r)
	})
}
