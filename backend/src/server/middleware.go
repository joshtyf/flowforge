package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	"github.com/joshtyf/flowforge/src/util"
	"github.com/joshtyf/flowforge/src/validation"
	"go.mongodb.org/mongo-driver/mongo"
)

func validateCreatePipelineRequest(next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pipeline, err := decode[models.PipelineModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, err)
			return
		}

		err = validation.ValidatePipeline(&pipeline)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to validate pipeline: %s", err))
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

func isAuthenticated(next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse the issuer url: %s", err))
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
			logger.Error(fmt.Sprintf("failed to set up jwt validator: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
			logger.Error(fmt.Sprintf("failed to validate jwt: %s", err))
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
func isAuthorisedAdmin(next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		claims := token.CustomClaims.(*CustomClaims)
		requiredPermission := "approve:pipeline_step"
		if !claims.HasPermission(requiredPermission) {
			logger.Error(fmt.Sprintf("unauthorized: missing permission %s", requiredPermission))
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isOrgOwner(postgresClient *sql.DB, next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(util.OrgContextKey{}).(int)
		membership, err := getMembership(org_id, postgresClient, r)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("user not authorized owner")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("unable to verify ownership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if membership.Role != models.Owner {
			logger.Error("user not authorized owner")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOrgAdmin(postgresClient *sql.DB, next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(util.OrgContextKey{}).(int)
		membership, err := getMembership(org_id, postgresClient, r)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("user not authorized admin")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("unable to verify admin role: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if membership.Role == models.Member {
			logger.Error("user not authorized admin")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOrgMember(postgresClient *sql.DB, next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(util.OrgContextKey{}).(int)
		_, err := getMembership(org_id, postgresClient, r)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("user not authorized member")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("unable to verify membership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getOrgIdFromQuery(next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var org_id int
		var err error
		if q := r.URL.Query().Get("org_id"); q == "" {
			logger.Error("org id does not exist in query")
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		} else {
			org_id, err = strconv.Atoi(q)
			if err != nil {
				logger.Error(fmt.Sprintf("failed to parse org id as integer: %s", err))
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
				return
			}
		}

		r = r.Clone(context.WithValue(r.Context(), util.OrgContextKey{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func getOrgIdFromPath(next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var org_id int
		var err error
		vars := mux.Vars(r)
		if id := vars["organizationId"]; id == "" {
			logger.Error("org id does not exist in path")
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		} else {
			org_id, err = strconv.Atoi(id)
			if err != nil {
				logger.Error(fmt.Sprintf("failed to parse org id as integer: %s", err))
				encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
				return
			}
		}

		r = r.Clone(context.WithValue(r.Context(), util.OrgContextKey{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func getOrgIdFromRequestBody(next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type OrgId struct {
			OrganizationId int `json:"org_id"`
		}
		org, err := decode[OrgId](r)
		if err != nil {
			logger.Error("unable to parse request body into json")
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
			return
		}
		org_id := org.OrganizationId

		if org_id == 0 {
			logger.Error("org id does not exist in request body")
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		}

		r = r.Clone(context.WithValue(r.Context(), util.OrgContextKey{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func getOrgIdUsingSrId(mongoClient *mongo.Client, next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var org_id int
		vars := mux.Vars(r)
		sr_id := vars["requestId"]
		if sr_id == "" {
			logger.Error("service request id does not exist in path")
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrUnauthorised, http.StatusBadRequest))
			return
		}
		sr, err := database.NewServiceRequest(mongoClient).GetById(sr_id)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to retrieve service request by service request id: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
			return
		}
		org_id = sr.OrganizationId

		r = r.Clone(context.WithValue(r.Context(), util.OrgContextKey{}, org_id))
		next.ServeHTTP(w, r)
	})
}

func validateMembershipChange(postgresClient *sql.DB, next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(util.OrgContextKey{}).(int)
		requestorMembership, err := getMembership(org_id, postgresClient, r)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to get requestor membership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
			return
		}

		if requestorMembership.Role == models.Member {
			logger.Error("user not authorized to grant/delete membership")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		targetMembership, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		if targetMembership.UserId == requestorMembership.UserId {
			logger.Error("user not authorized to change own membership")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		err = models.ValidateRole(targetMembership.Role)
		if err != nil {
			logger.Error("unable to add/update membership as role is invalid")
			encode(w, r, http.StatusBadRequest, newHandlerError(err, http.StatusBadRequest))
			return
		}

		if targetMembership.Role == models.Owner {
			logger.Error("unable to grant/delete ownership")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnableModifyOwnership, http.StatusForbidden))
			return
		}

		targetExistingMembership, err := database.NewMembership(postgresClient).GetMembershipByUserAndOrgId(targetMembership.UserId, targetMembership.OrgId)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			logger.Error(fmt.Sprintf("unable to verify subject role: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		if targetExistingMembership != nil && targetExistingMembership.Role == models.Owner {
			logger.Error("user not authorized to alter/delete ownership")
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

func validateOwnershipTransfer(postgresClient *sql.DB, next http.Handler, logger logger.ServerLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		org_id := r.Context().Value(util.OrgContextKey{}).(int)
		requestorMembership, err := getMembership(org_id, postgresClient, r)
		if err != nil {
			logger.Error(fmt.Sprintf("unable to get requestor membership: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrUnauthorised, http.StatusInternalServerError))
			return
		}

		if requestorMembership.Role != models.Owner {
			logger.Error("user not authorized to transfer ownership")
			encode(w, r, http.StatusForbidden, newHandlerError(ErrUnauthorised, http.StatusForbidden))
			return
		}

		targetMembership, err := decode[models.MembershipModel](r)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse json request body: %s", err))
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrJsonParseError, http.StatusBadRequest))
			return
		}

		_, err = database.NewMembership(postgresClient).GetMembershipByUserAndOrgId(targetMembership.UserId, targetMembership.OrgId)
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("target user for transfer is not part of organisation")
			encode(w, r, http.StatusBadRequest, newHandlerError(ErrNotOrgMember, http.StatusBadRequest))
			return
		} else if err != nil {
			logger.Error(fmt.Sprintf("unable to verify target role: %s", err))
			encode(w, r, http.StatusInternalServerError, newHandlerError(ErrInternalServerError, http.StatusInternalServerError))
			return
		}

		next.ServeHTTP(w, r)
	})
}
