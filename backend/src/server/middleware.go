package server

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/validation"
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

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Permissions string `json:"permissions"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
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

func isAuthenticated(next http.Handler) http.Handler {
	// TODO: implement a proper flag pattern
	env := os.Getenv("ENV")
	if env == "dev" {
		logger.Info("[Authentication] Skipping authentication in dev environment", nil)
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
		if err != nil {
			logger.Error("[Authentication] Failed to parse the issuer url", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{os.Getenv("AUTH0_AUDIENCE")},
			validator.WithCustomClaims(
				func() validator.CustomClaims {
					return &CustomClaims{}
				},
			),
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			logger.Error("[Authentication] Failed to set up jwt validator", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			logger.Error("[Authentication] Encountered error while validating JWT", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnableToValidateJWT, http.StatusUnauthorized))
		}

		middleware := jwtmiddleware.New(
			jwtValidator.ValidateToken,
			jwtmiddleware.WithErrorHandler(errorHandler),
		)
		middleware.CheckJWT(next).ServeHTTP(w, r)
	})

}

// TODO: To be tested once frontend is ready
func isAuthorisedAdmin(next http.Handler) http.Handler {
	env := os.Getenv("ENV")
	if env == "dev" {
		logger.Info("[Authorization] Skipping admin check in dev environment", nil)
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*CustomClaims)
		if !claims.HasPermission("approve:pipeline_step") {
			logger.Error("[Authorization] User not authorized admin", nil)
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}
		next.ServeHTTP(w, r)
	})
}
