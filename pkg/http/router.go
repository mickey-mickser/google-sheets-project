package http

import (
	"github.com/go-chi/chi"
	chimiddlware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/mickey-mickser/telegram-project/pkg/http/handler"
	usecase "github.com/mickey-mickser/telegram-project/pkg/storage/sheets"
	"github.com/sirupsen/logrus"
	"net/http"
)

func newRouter(log logrus.FieldLogger, sheetUse usecase.SheetUseCase) http.Handler {
	m := chi.NewMux()
	m.Use(
		Logger("router", log),
		chimiddlware.Recoverer,
		cors.AllowAll().Handler,
	)

	h := handler.NewHandler(log, sheetUse)
	m.Route("/sheets", func(r chi.Router) {
		r.Get("/", h.Get)
		//r.Post("/", h.Post)
		//r.Put("/", h.Put)
		//r.Delete("/", h.Delete)

	})
	m.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return m
}
