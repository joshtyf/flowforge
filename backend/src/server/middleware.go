package server

import (
	"net/http"

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
			encode(w, r, http.StatusBadRequest, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
