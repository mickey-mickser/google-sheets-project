package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	srv *http.Server
	log logrus.FieldLogger
}

func NewServer(log logrus.Ext1FieldLogger, port string, handler *gin.Engine) *Server {
	return &Server{
		srv: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20, // 1 mb
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		log: log.WithField("runner", "server"),
	}
}

func (s *Server) Run(ctx context.Context, group *sync.WaitGroup) {
	group.Add(1)
	go func() {
		defer group.Done()
		s.monitorShutdown(ctx)
	}()
	go s.listen()
}

func (s *Server) listen() {
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.Error(err)
	} else {
		s.log.Info("server is stopped gracefully")
	}
}

// Monitor shutdown

func (server *Server) monitorShutdown(ctx context.Context) {
	select {
	case <-ctx.Done(): // Wait for interrupt signal
		if err := server.srv.Shutdown(ctx); err != nil {
			server.log.Error(errors.Wrap(err, "server shutdown error"))
		}
	}
}
