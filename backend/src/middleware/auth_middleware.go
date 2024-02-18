package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"github.com/joshtyf/flowforge/src/api"
	"github.com/joshtyf/flowforge/src/authenticator"
	"github.com/joshtyf/flowforge/src/logger"
	"golang.org/x/oauth2"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		accessToken := r.Header.Get("Authorization")
		session, _ := store.Get(r, accessToken)

		var profile map[string]interface{}
		profile, ok := session.Values["profile"].(map[string]interface{})

		if ok {
			logger.Info("[Authorization] Previously authorised", nil)
			next.ServeHTTP(w, r)
		}

		authToken := &oauth2.Token{AccessToken: accessToken}
		if accessToken == "" {
			authError := &api.HandlerError{Error: errors.New("authorization token missing"), Code: http.StatusUnauthorized}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		auth, err := authenticator.New()
		if err != nil {
			logger.Error("[Authorization] Failed to initialize authenticator", map[string]interface{}{"err": err})
			authError := &api.HandlerError{Error: api.ErrInternalServerError, Code: http.StatusInternalServerError}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		idToken, err := auth.VerifyIDToken(r.Context(), authToken)
		if err != nil {
			logger.Error("[Authorization] Unable to verify token", map[string]interface{}{"err": err})
			authError := &api.HandlerError{Error: errors.New("unable to verify token"), Code: http.StatusUnauthorized}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		err = idToken.Claims(&profile)
		if err != nil {
			logger.Error("[Authorization] Unable retrieve profile", map[string]interface{}{"err": err})
			authError := &api.HandlerError{Error: errors.New("unable to retrieve profile"), Code: http.StatusInternalServerError}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		session.Values["access_token"] = authToken.AccessToken
		session.Values["profile"] = profile
		err = session.Save(r, w)
		if err != nil {
			logger.Error("[Authorization] Unable save session", map[string]interface{}{"err": err})
			authError := &api.HandlerError{Error: errors.New("unable to save session"), Code: http.StatusInternalServerError}
			w.WriteHeader(authError.Code)
			json.NewEncoder(w).Encode(authError)
			return
		}

		// Call the next middleware function or final handler
		next.ServeHTTP(w, r)
	})
}
