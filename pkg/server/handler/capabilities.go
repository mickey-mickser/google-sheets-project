package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
)

// Capabilities reads ./capabilities.json and returns it as JSON. reads
// @Summary Get capabilities
// @Description Reads ./capabilities.json and returns it parsed as JSON
// @Tags Utils
// @Produce json
// @Success 200 {object} object "Raw JSON"
// @Failure 404 {object} responses.ResponseError "File not found"
// @Failure 400 {object} responses.ResponseError "Invalid JSON"
// @Failure 500 {object} responses.ResponseError "Internal error"
// @Router /api/v1/sheets/capabilities [get]
func (h *Handler) Capabilities(c *gin.Context) {
	b, err := os.ReadFile("./capabilities.json")
	if err != nil {
		if os.IsNotExist(err) {
			responses.NotFound(c, responses.NewNotFoundError("file not found"))
			return
		}
		responses.InternalError(c)
		return
	}

	var payload any
	if err := json.Unmarshal(b, &payload); err != nil {
		responses.BadRequest(c, responses.NewInvalidBodyError(errors.New("file is not valid JSON")))
		return
	}

	c.JSON(http.StatusOK, payload)
}
