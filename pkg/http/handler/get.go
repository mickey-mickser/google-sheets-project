package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Get handles the request to get data from a Google Sheets document
// @Summary Get data from Google Sheets
// @Description This endpoint retrieves a specified range of data from a Google Sheets document
// @Tags Sheets
// @Accept json
// @Produce json
// @Param sheet_id query string true "Sheet ID"
// @Param range query string false "Range"
// @Success 200 {object} map[string]interface{} "Successfully retrieved data"
// @Failure 400 {object} map[string]interface{} "Missing Sheet ID or invalid parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /read [get]
func (h *Handler) Get(c *gin.Context) {

	rangeParam := c.Query("range")
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		errMsg := "Sheet ID parameter is required"
		h.Log.WithFields(logrus.Fields{
			"handler": "Get",
		}).Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	resp, err := h.SheetUse.Get(c.Request.Context(), sheetId, rangeParam)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get from Google Sheets. Error: %v", err)
		h.Log.WithFields(logrus.Fields{
			"handler": "Get",
			"error":   err.Error(),
		}).Error(errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   resp,
	})
}
