// file: http/handler/echo.go
package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
	"github.com/sirupsen/logrus"
)

// EchoJSON reads raw JSON and returns it unchanged.
// @Summary Echo JSON
// @Description Reads JSON body and returns it unchanged (validates JSON).
// @Tags Utils
// @Accept json
// @Produce json
// @Param request body object true "Any JSON payload"
// @Success 200 {object} map[string]interface{} "Echoed JSON"
// @Failure 400 {object} responses.ResponseError "Invalid JSON"
// @Router /api/v1/utils/echo [post]
func (h *Handler) EchoJSON(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerCapabilities,
			"error":   err.Error(),
		}).Error(errGetCapabilities)
		responses.BadRequest(c, responses.NewInvalidBodyError(err))
		return
	}

	if !json.Valid(body) {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerCapabilities,
		}).Warn("invalid JSON")
		responses.BadRequest(c, responses.NewInvalidBodyError(errors.New("invalid JSON")))
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", body)
}
