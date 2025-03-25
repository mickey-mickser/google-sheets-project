package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BatchUpdateRequest struct {
	Requests []map[string]interface{} `json:"requests"`
}

func (h *Handler) BatchUpdate(c *gin.Context) {
	sheetId := c.Query("sheet_id")
	if sheetId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Table parameter is required"})
		return
	}
	var req BatchUpdateRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.sheetUse.BatchUpdate(c.Request.Context(), sheetId, req.Requests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   res,
	})
}
