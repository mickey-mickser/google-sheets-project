package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
	"github.com/sirupsen/logrus"
)

// BatchUpdate updates values in a Google Sheet.
// @Summary Batch update Google Sheets
// @Description Accepts a list of updates and applies them to the specified Google Sheet.
// @Tags Sheets
// @Accept json
// @Produce json
// @Param sheet_id query string true "Google Sheet ID"
// @Param request body models.BatchUpdateRequest true "List of updates..."
// @Success 200 {object} responses.BatchUpdateValuesResponse "Success response with status and data"
// @Failure 400 {object} responses.ResponseError "Bad request with error message"
// @Failure 500 {object} responses.ResponseError "Internal server error"
// @Router /api/v1/sheets/update [post]
func (h *Handler) BatchUpdate(c *gin.Context) {
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		h.Log.WithField("handler", handlerNameBatchUpdate).Error(errMissingSheetID)
		responses.BadRequest(c, responses.NewMissingParamError("sheet_id"))
		return
	}

	var req models.BatchUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {

		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameBatchUpdate,
			"error":   err,
		}).Error("Failed to bind JSON")
		responses.BadRequest(c, responses.NewBadRequestError(err))
		return
	}

	resp, err := h.SheetUse.BatchUpdate(c.Request.Context(), sheetId, req.Updates)
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameBatchUpdate,
			"error":   err,
		}).Error(errUpdateFailed)
		responses.InternalError(c)
		return
	}
	responses.OK(c, gin.H{"status": "ok", "data": resp})

}
