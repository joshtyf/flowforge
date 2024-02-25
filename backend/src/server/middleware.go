package server

import (
	"encoding/json"
	"net/http"

	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
	"github.com/joshtyf/flowforge/src/validation"
)

type customMiddleWareFunc func(customHandlerFunc) customHandlerFunc

func ValidateCreatePipelineRequest(nextHandlerFunc customHandlerFunc) customHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) *HandlerError {
		pipeline := &models.PipelineModel{}
		json.NewDecoder(r.Body).Decode(pipeline)
		err := validation.ValidatePipeline(pipeline)
		if err != nil {
			logger.Error("[ValidateCreatePipelineRequest] Error validating pipeline", map[string]interface{}{"err": err})
			return NewHandlerError(err, http.StatusBadRequest)
		}

		return nextHandlerFunc(w, r)
	}
}
