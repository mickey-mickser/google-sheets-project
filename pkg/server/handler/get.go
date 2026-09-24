package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mickey-mickser/google-sheets-project/pkg/server/responses"
	"github.com/sirupsen/logrus"
	"regexp"
	"time"
)

// Get handles the request to get data from a Google Sheets document
// @Summary Get data from Google Sheets
// @Description This endpoint retrieves a specified range of data from a Google Sheets document
// @Tags Sheets
// @Accept json
// @Produce json
// @Param sheet_id query string true "Sheet ID"
// @Param range query string false "Range"
// @Success 200 {object} [][]interface{} "Successfully retrieved data"
// @Failure 400 {object} responses.ResponseError "Missing Sheet ID or invalid parameters"
// @Failure 500 {object} responses.ResponseError "Internal server error"
// @Router /api/v1/sheets/read [get]
func (h *Handler) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		h.Log.WithField("handler", handlerNameGet).Error(errMissingSheetID)
		responses.BadRequest(c, responses.NewMissingParamError("sheet_id"))
		return
	}
	// validate range format (accepts "A1" or "A1:B10")
	rangeParam := c.Query("range")
	if err := validateRange(rangeParam); err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameGet,
			"range":   rangeParam,
		}).Error("invalid range format")
		responses.BadRequest(c, responses.NewBadRequestError(fmt.Errorf("invalid range format")))
		return
	}
	resp, err := h.SheetUse.Get(ctx, sheetId, rangeParam)
	if err != nil {
		h.Log.WithFields(logrus.Fields{
			"handler": handlerNameGet,
			"error":   err,
		}).Error(errGetFailed)
		responses.InternalError(c)
		return
	}

	responses.OK(c, gin.H{"status": "success", "data": resp})
}
func validateRange(rangeParam string) error {
	if rangeParam != "" && !regexp.MustCompile(`^[A-Z]+[0-9]+(:[A-Z]+[0-9]+)?$`).MatchString(rangeParam) {
		return fmt.Errorf("invalid range format")
	}
	return nil
}
