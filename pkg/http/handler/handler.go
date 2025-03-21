package handler

import (
	usecase "github.com/mickey-mickser/telegram-project/pkg/storage/sheets"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	log      logrus.FieldLogger
	sheetUse usecase.SheetUseCase
}

func NewHandler(log logrus.FieldLogger, sheetUse usecase.SheetUseCase) Handler {
	return &handler{
		log:      log,
		sheetUse: sheetUse,
	}
}
