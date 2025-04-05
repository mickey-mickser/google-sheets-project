package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// Create handles the request to create a new table in Google Sheets
// @Summary Create a new table in Google Sheets
// @Description This endpoint creates a new table in Google Sheets based on the provided request body
// @Tags Sheets
// @Accept json
// @Produce json
// @Param body body string true "Table Data"
// @Success 200 {object} map[string]interface{} "Successfully created table"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /create [post]
func (h *Handler) Create(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errMsg := "Invalid request body"
		h.Log.WithFields(logrus.Fields{
			"handler": "CreateTable",
			"error":   err.Error(),
		}).Error(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	resp, err := h.SheetUse.Create(c.Request.Context(), body)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create tables in Google Sheets. Error: %v", err)
		h.Log.WithFields(logrus.Fields{
			"handler": "CreateTable",
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
