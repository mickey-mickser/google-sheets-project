package handler

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/mickey-mickser/telegram-project/pkg/storage/sheets"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	sheetUse usecase.SheetUseCase
	log      logrus.FieldLogger
}

func NewHandler(log logrus.FieldLogger, sheetUse usecase.SheetUseCase) *Handler {
	return &Handler{
		sheetUse: sheetUse,
		log:      log,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	sheets := router.Group("/sheets")
	{
		sheets.POST("", h.BatchUpdate)
	}

	return router
}
