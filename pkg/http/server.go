package http

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type Http struct {
	srv *http.Server
	log logrus.FieldLogger
}

func NewHttp(log logrus.Ext1FieldLogger, handler http.Handler) *Http {
	return &Http{
		srv: &http.Server{
			Addr:    ":8099",
			Handler: handler,
		},
		log: log.WithField("runner", "server"),
	}
}

func (a *Http) Run(ctx context.Context, group *sync.WaitGroup) {
	group.Add(1)
	go func() {
		defer group.Done()
		a.monitorShutdown(ctx)
	}()
	go a.listen()
}

func (a *Http) listen() {
	if err := a.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		a.log.Error(err)
	} else {
		a.log.Info("http server is stopped gracefully")
	}
}

func (http *Http) monitorShutdown(ctx context.Context) {
	select {
	case <-ctx.Done():
		if err := http.srv.Shutdown(ctx); err != nil {
			http.log.Error(errors.Wrap(err, "http server shutdown error"))
		}
	}
}

//package http
//
//import (
//	"context"
//	"net/http"
//	"time"
//)
//
//type Server struct {
//	httpServer *http.Server
//}
//
//func (s *Server) Run(port string, handler http.Handler) error {
//
//	s.httpServer = &http.Server{
//		Addr:           ":" + port,
//		Handler:        handler,
//		MaxHeaderBytes: 1 << 20, // 1 mb
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//	}
//	return s.httpServer.ListenAndServe()
//}
//
//func (s *Server) Shutdown(ctx context.Context) error {
//	return s.httpServer.Shutdown(ctx)
//}
