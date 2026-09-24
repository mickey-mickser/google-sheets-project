package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
	"github.com/sirupsen/logrus"
)

const theWholeSheet = "A1:Z1000"

// ClearTable handles the request to delete data from a Google Sheets document
// @Summary ClearTable data from Google Sheets
// @Description This endpoint clear table in a Google Sheets document
// @Tags Sheets
// @Accept json
// @Produce json
// @Param sheet_id query string true "Sheet ID"
// @Param range query string false "Range"
// @Param request body models.DeleteRequest true "Delete request payload"
// @Success 200 {object} responses.ClearValuesResponse "Successfully clear table"
// @Failure 400 {object} responses.ResponseError "Missing required parameters"
// @Failure 500 {object} responses.ResponseError "Internal server error"
// @Router /api/v1/sheets/clearTable [post]
func (h *Handler) ClearTable(c *gin.Context) {
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		h.Log.WithField("handler", handlerNameClearTable).Error(errMissingSheetID)
		responses.BadRequest(c, responses.NewMissingParamError("sheet_id"))
		return
	}

	// Create an instance of DeleteRequest (your struct to bind the body)
	var req models.DeleteRequest
	// Use ShouldBindJSON to bind the body to the DeleteRequest struct
	if err := c.ShouldBindJSON(&req); err != nil {
		errMsg := errInvalidBody
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameClearTable,
			"error":   err.Error(),
		}).Error(errMsg)
		responses.BadRequest(c, responses.NewInvalidBodyError(err))
		return
	}

	//If the value is empty, then we select the entire table range
	if req.Range == "" {
		req.Range = theWholeSheet
	}

	// After binding, you can now safely use deleteRequest
	resp, err := h.SheetUse.Delete(c.Request.Context(), sheetId, req.Range)
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameClearTable,
			"error":   err.Error(),
		}).Error(errDeleteFailed)
		responses.InternalError(c)
		return
	}

	responses.OK(c, gin.H{"status": "ok", "data": resp})
}
