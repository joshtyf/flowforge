package server

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/validation"
	"go.mongodb.org/mongo-driver/mongo"
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
	Permissions []string `json:"permissions"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// HasPermission checks whether our claims have a specific permission.
// In our case, since we are using this to check if user is admin, will be checking for approve:pipeline_step permission
func (c CustomClaims) HasPermission(expectedPermission string) bool {
	for i := range c.Permissions {
		if c.Permissions[i] == expectedPermission {
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
			encode(w, r, http.StatusUnauthorized, newHandlerError(ErrUnableToValidateJWT, http.StatusUnauthorized))
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
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}
		next.ServeHTTP(w, r)
	})
}

type OrgId struct {
	OrganisationId int `json:"org_id"`
}

func isOrgOwner(postgresClient *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(OrgId{}).(int)
		membership, err := getMembership(org_id, postgresClient, r)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("[Authorization] User not authorized owner", nil)
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		} else if err != nil {
			logger.Error("[Authorization] Error encountered while verifying ownership", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if membership.Role != models.Owner {
			logger.Error("[Authorization] User not authorized owner", nil)
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOrgAdmin(postgresClient *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(OrgId{}).(int)
		membership, err := getMembership(org_id, postgresClient, r)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("[Authorization] User not authorized member", nil)
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		} else if err != nil {
			logger.Error("[Authorization] Error encountered while verifying admin role", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if membership.Role == models.Member {
			logger.Error("[Authorization] User not authorized admin", nil)
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOrgMember(postgresClient *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(OrgId{}).(int)
		_, err := getMembership(org_id, postgresClient, r)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("[Authorization] User not authorized member", nil)
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		} else if err != nil {
			logger.Error("[Authorization] Error encountered while verifying membership", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getOrgIdFromQuery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var org_id int
		var err error
		if q := r.URL.Query().Get("org_id"); q == "" {
			logger.Error("[Authorization] Org Id does not exist in query", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		} else {
			org_id, err = strconv.Atoi(q)
			if err != nil {
				logger.Error("[Authorization] Failed to parse Org Id as integer", nil)
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
				return
			}
		}

		r = r.Clone(context.WithValue(r.Context(), OrgId{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func getOrgIdFromPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var org_id int
		var err error
		vars := mux.Vars(r)
		if id := vars["organisationId"]; id == "" {
			logger.Error("[Authorization] Org Id does not exist in path", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		} else {
			org_id, err = strconv.Atoi(id)
			if err != nil {
				logger.Error("[Authorization] Failed to parse Org Id as integer", nil)
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
				return
			}
		}

		r = r.Clone(context.WithValue(r.Context(), OrgId{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func getOrgIdFromRequestBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org, err := decode[OrgId](r)
		if err != nil {
			logger.Error("[Authorization] Unable to parse request body into json", nil)
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
			return
		}
		org_id := org.OrganisationId

		if org_id == 0 {
			logger.Error("[Authorization] Org Id does not exist in request body", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		}

		r = r.Clone(context.WithValue(r.Context(), OrgId{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func getOrgIdUsingSrId(mongoClient *mongo.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var org_id int
		vars := mux.Vars(r)
		if sr_id := vars["requestId"]; sr_id == "" {
			logger.Error("[Authorization] Service Request Id does not exist in path", nil)
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		} else {
			sr, err := database.NewServiceRequest(mongoClient).GetById(sr_id)
			if err != nil {
				logger.Error("[Authorization] Failed to retrieve service request by service request id", nil)
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
				return
			}

			org_id = sr.OrganisationId
		}

		r = r.Clone(context.WithValue(r.Context(), OrgId{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func validateMembershipChange(postgresClient *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(OrgId{}).(int)
		membership, err := getMembership(org_id, postgresClient, r)
		if err != nil {
			logger.Error("[Authorization] Error getting requestor membership", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
			return
		}

		mm, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error("[Authorization] Unable to parse json request body", map[string]interface{}{"err": err})
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		if mm.Role == models.Owner && membership.Role == models.Admin {
			logger.Error("[Authorization] User not authorized to grant/delete ownership", nil)
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		subjectMembership, err := database.NewMembership(postgresClient).GetMembershipByUserAndOrgId(mm.UserId, mm.OrgId)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			logger.Error("[Authorization] Error encountered while verifying subject role", map[string]interface{}{"err": err})
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if subjectMembership != nil && subjectMembership.Role == models.Owner && membership.Role == models.Admin {
			logger.Error("[Authorization] User not authorized to alter/delete ownership", nil)
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getMembership(org_id int, postgresClient *sql.DB, r *http.Request) (*models.MembershipModel, error) {
	token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	userId := token.RegisteredClaims.Subject

	mm, err := database.NewMembership(postgresClient).GetMembershipByUserAndOrgId(userId, org_id)
	if err != nil {
		return nil, err
	}

	return mm, nil
}
