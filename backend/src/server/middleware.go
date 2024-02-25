package server

import (
	"net/http"

	"github.com/joshtyf/flowforge/src/authenticator"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/validation"
	"golang.org/x/oauth2"
)

func validateCreatePipelineRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pipeline, err := decode[models.PipelineModel](r)
		if err != nil {
			logger.Error("[ValidateCreatePipelineRequest] Error decoding pipeline from request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, err)
			return
		}

		err = validation.ValidatePipeline(&pipeline)
		if err != nil {
			logger.Error("[ValidateCreatePipelineRequest] Error validating pipeline", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(err, http.StatusBadRequest))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		accessToken := r.Header.Get("Authorization")

		authToken := &oauth2.Token{AccessToken: accessToken}
		if accessToken == "" {
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrBearerTokenNotFound, http.StatusUnauthorized))
			return
		}

		auth, err := authenticator.New()
		if err != nil {
			logger.Error("[Authorization] Failed to initialize authenticator", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		idToken, err := auth.VerifyIDToken(r.Context(), authToken)
		if err != nil {
			logger.Error("[Authorization] Unable to verify token", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnableToVerifyBearerToken, http.StatusUnauthorized))
			return
		}

		// TODO: Review to see if this fits flow once frontend side of the flow is finalised
		var profile map[string]interface{}
		err = idToken.Claims(&profile)
		if err != nil {
			logger.Error("[Authorization] Unable retrieve profile", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnableToRetrieveProfile, http.StatusInternalServerError))
			return
		}

		// Call the next middleware function or final handler
		next.ServeHTTP(w, r)
	})
}
