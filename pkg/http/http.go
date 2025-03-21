package http

import (
	"context"
	usecase "github.com/mickey-mickser/telegram-project/pkg/storage/sheets"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type Http struct {
	srv *http.Server
	log logrus.FieldLogger
}

func NewHttp(log logrus.Ext1FieldLogger, sheetUse usecase.SheetUseCase) *Http {
	return &Http{
		srv: &http.Server{
			Addr:    ":8099",
			Handler: newRouter(log, sheetUse),
		},
		log: log.WithField("runner", "server"),
	}
}

func (h *Http) Run(ctx context.Context, group *sync.WaitGroup) {
	group.Add(1)
	go func() {
		defer group.Done()
		h.monitorShutdown(ctx)
	}()
	go h.listen()
}

func (h *Http) listen() {
	h.log.Infof("Starting HTTP server on port %s...", h.srv.Addr)

	if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		h.log.Error(err)
	} else {
		h.log.Info("http server is stopped gracefully")
	}
}

func (h *Http) monitorShutdown(ctx context.Context) {
	select {
	case <-ctx.Done():
		if err := h.srv.Shutdown(ctx); err != nil {
			h.log.Error(errors.Wrap(err, "http server shutdown error"))
		}
	}
}
