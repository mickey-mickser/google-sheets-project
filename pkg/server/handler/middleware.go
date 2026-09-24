package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (h *Handler) loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		h.Log.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"ip":     c.ClientIP(),
		}).Info("incoming request")

		c.Next()

		duration := time.Since(start)
		h.Log.WithFields(logrus.Fields{
			"status":   c.Writer.Status(),
			"duration": duration,
		}).Info("request completed")
	}
}

// corsMiddleware sets CORS headers based on AllowedOrigins
func (h *Handler) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		// default to wildcard
		allowed := "*"
		for _, o := range h.AllowedOrigins {
			if o == "*" || o == origin {
				allowed = o
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowed)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// Health check endpoint
func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up", "service": "google-sheets-service"})
}
