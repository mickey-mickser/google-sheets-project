package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// BatchUpdate updates values in a Google Sheet.
//
// @Summary Batch update Google Sheets
// @Description Accepts a list of updates and applies them to the specified Google Sheet.
// @Tags Sheets
// @Accept json
// @Produce json
// @Param sheet_id query string true "Google Sheet ID"
// @Param updates body []map[string]string true "List of updates, each with 'range' and 'value'"
// @Success 200 {object} map[string]interface{} "Success response with status and data"
// @Failure 400 {object} map[string]string "Bad request with error message"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /update [post]
func (h *Handler) BatchUpdate(c *gin.Context) {
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		errMsg := "Sheet ID parameter is required"
		h.Log.WithFields(logrus.Fields{
			"handler": "BatchUpdate",
		}).Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errMsg := "Invalid request body"
		h.Log.WithFields(logrus.Fields{
			"handler": "BatchUpdate",
			"error":   err.Error(),
		}).Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	resp, err := h.SheetUse.BatchUpdate(c.Request.Context(), sheetId, body)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to update the Google Sheets. Error: %v", err)
		h.Log.WithFields(logrus.Fields{
			"handler": "BatchUpdate",
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
