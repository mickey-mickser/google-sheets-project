package handler

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/mickey-mickser/google-sheets-project/pkg/usecase/sheets"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	SheetUse usecase.SheetUseCase
	Log      logrus.FieldLogger
}

func NewHandler(log logrus.FieldLogger, sheetUse usecase.SheetUseCase) *Handler {
	return &Handler{
		SheetUse: sheetUse,
		Log:      log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	sheets := router.Group("/sheets")
	{
		sheets.POST("/update", h.BatchUpdate)
		sheets.GET("/read", h.Get)
		sheets.POST("/create", h.Create)
		sheets.POST("/delete", h.Delete)

	}

	return router
}
