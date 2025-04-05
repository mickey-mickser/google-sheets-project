package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// Delete handles the request to delete data from a Google Sheets document
// @Summary Delete data from Google Sheets
// @Description This endpoint deletes data from a specified range in a Google Sheets document
// @Tags Sheets
// @Accept json
// @Produce json
// @Param sheet_id query string true "Sheet ID"
// @Param range query string false "Range"
// @Param body body string true "Table Data"
// @Success 200 {object} map[string]interface{} "Successfully deleted data"
// @Failure 400 {object} map[string]interface{} "Missing required parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /delete [delete]
func (h *Handler) Delete(c *gin.Context) {
	rangeParam := c.Query("range")
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		errMsg := "Sheet ID parameter is required"
		h.Log.WithFields(logrus.Fields{
			"handler": "ClearTable",
		}).Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errMsg := "Invalid request body"
		h.Log.WithFields(logrus.Fields{
			"handler": "ClearTable",
			"error":   err.Error(),
		}).Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	resp, err := h.SheetUse.Delete(c.Request.Context(), sheetId, rangeParam, body)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create tables in Google Sheets. Error: %v", err)
		h.Log.WithFields(logrus.Fields{
			"handler": "ClearTable",
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
