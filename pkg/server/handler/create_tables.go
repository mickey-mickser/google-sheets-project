package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/models"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
	"github.com/sirupsen/logrus"
)

// Create handles the request to create a new table in Google Sheets
// @Summary Create a new table in Google Sheets
// @Description This endpoint creates a new table in Google Sheets based on the provided request body
// @Tags Sheets
// @Accept json
// @Produce json
// @Param request body models.CreateRequest true "Table Data"
// @Success 200 {object} responses.Spreadsheet "Successfully created table"
// @Failure 400 {object} responses.ResponseError "Invalid request body"
// @Failure 500 {object} responses.ResponseError "Internal server error"
// @Router /api/v1/sheets/create [post]
func (h *Handler) Create(c *gin.Context) {
	var req models.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameCreateTable,
			"error":   err.Error(),
		}).Error("Failed to bind JSON")
		responses.BadRequest(c, responses.NewInvalidBodyError(err))
		return
	}

	h.Log.WithFields(logrus.Fields{
		"handler": handlerNameCreateTable,
		"title":   req.Properties.Title,
	}).Info("creating new sheet")

	resp, err := h.SheetUse.Create(c.Request.Context(), req)
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameCreateTable,
			"error":   err.Error(),
		}).Error(errCreateFailed)
		responses.InternalError(c)
		return
	}

	responses.OK(c, gin.H{"status": "ok", "data": resp})
}
