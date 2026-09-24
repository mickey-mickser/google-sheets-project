package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
	"github.com/sirupsen/logrus"
)

// BatchUpdate updates values in a Google Sheet with explicit error details on failure.
func (h *Handler) BatchUpdate(c *gin.Context) {
	// Retrieve sheet ID from query
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		h.Log.WithField("handler", handlerNameBatchUpdate).Error(errMissingSheetID)
		responses.BadRequest(c, responses.NewMissingParamError("sheet_id"))
		return
	}

	// Bind JSON request
	var req models.BatchUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameBatchUpdate,
			"error":   err,
		}).Error("Failed to bind JSON")
		responses.BadRequest(c, responses.NewBadRequestError(err))
		return
	}

	// Execute batch update
	resp, err := h.SheetUse.BatchUpdate(c.Request.Context(), sheetId, req.Updates)
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameBatchUpdate,
			"error":   err,
		}).Error(errUpdateFailed)

		// Return explicit error details in the response
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": errUpdateFailed,
		})
		return
	}

	// Success response
	responses.OK(c, gin.H{"status": "ok", "data": resp})
}
