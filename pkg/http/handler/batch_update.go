package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

type BatchUpdateRequest struct {
	Requests []map[string]interface{} `json:"requests"`
}

func (h *Handler) BatchUpdate(c *gin.Context) {
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		errMsg := "Table parameter is required"
		log.Println("ERROR:", errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errMsg := "Invalid request body"
		log.Println("ERROR:", errMsg, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	res, err := h.sheetUse.BatchUpdate(c.Request.Context(), sheetId, body)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to update the Google Sheet. Error: %v", err)
		log.Println("ERROR:", errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   res,
	})
}
