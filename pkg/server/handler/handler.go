package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mickey-mickser/google-sheets-project/cmd/docs"
	"github.com/mickey-mickser/google-sheets-project/pkg/config"
	"github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/sheets/v4"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Add these constants at the package level
const (
	handlerNameBatchUpdate = "BatchUpdate"
	handlerNameClearTable  = "ClearTable"
	handlerNameCreateTable = "CreateTable"
	handlerNameGet         = "Get"
	handlerNameShare       = "Share"
	handlerCapabilities    = "Capabilities"

	errMissingSheetID  = "Sheet ID parameter is required"
	errInvalidBody     = "Invalid request body"
	errUpdateFailed    = "Failed to update the Google Sheets"
	errDeleteFailed    = "Failed to delete data from Google Sheets"
	errCreateFailed    = "Failed to create tables in Google Sheets"
	errGetFailed       = "Failed to get from Google Sheets"
	errShareFailed     = "Failed to share permission foe reading Google Sheets"
	errGetCapabilities = "Failed to get Capabilities from Google Sheets UseCase"
)

type Handler struct {
	SheetUse       sheetUsecase.SheetUseCase
	Log            logrus.FieldLogger
	AllowedOrigins []string // allowed CORS origins, from ALLOWED_ORIGINS env var
}

// NewHandler initializes the HTTP handler with dependencies and CORS settings
func NewHandler(log *logrus.Logger, sheetSrv *sheets.Service, driveSvc *drive.Service, permission *config.PermissionsStruct, sheetID string) *Handler {
	// Read allowed origins from environment variable, comma-separated; default to wildcard
	origEnv := os.Getenv("ALLOWED_ORIGINS")
	var allowed []string
	if origEnv == "" {
		allowed = []string{"*"}
	} else {
		for _, o := range strings.Split(origEnv, ",") {
			trimmed := strings.TrimSpace(o)
			if trimmed != "" {
				allowed = append(allowed, trimmed)
			}
		}
	}

	sheetUse := sheetUsecase.NewSheetUse(sheetSrv, driveSvc, log, permission, sheetID)

	return &Handler{
		SheetUse:       sheetUse,
		Log:            log,
		AllowedOrigins: allowed,
	}
}

// InitRoutes sets up the Gin engine, middleware, and routes
func (h *Handler) InitRoutes() *gin.Engine {
	// Running in "debug" mode. Switch to "release" mode in production.
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(h.loggerMiddleware())
	// global timeout for all requests
	router.Use(func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})
	router.Use(h.corsMiddleware())

	router.GET("/health", h.healthCheck)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/debug/metadata/email", func(c *gin.Context) {
		req, _ := http.NewRequest("GET",
			"http://metadata/computeMetadata/v1/instance/service-accounts/default/email", nil)
		req.Header.Add("Metadata-Flavor", "Google")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		c.String(resp.StatusCode, string(body))
	})

	router.GET("/debug/metadata/token", func(c *gin.Context) {
		req, _ := http.NewRequest("GET",
			"http://metadata/computeMetadata/v1/instance/service-accounts/default/token", nil)
		req.Header.Add("Metadata-Flavor", "Google")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		c.Data(resp.StatusCode, "application/json", body)
	})

	api := router.Group("/api/v1")
	{
		s := api.Group("/sheets")
		{
			s.POST("/share", h.Share)
			s.POST("/update", h.BatchUpdate)
			s.GET("/read", h.Get)
			s.POST("/create", h.Create)
			s.POST("/clearTable", h.ClearTable)
			s.GET("/capabilities", h.Capabilities)
		}
	}
	return router
}
