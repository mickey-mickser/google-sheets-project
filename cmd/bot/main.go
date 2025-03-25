package main

import (
	"context"
	"github.com/mickey-mickser/telegram-project/pkg/clients/sheets"
	"github.com/mickey-mickser/telegram-project/pkg/config"
	"github.com/mickey-mickser/telegram-project/pkg/http"
	"github.com/mickey-mickser/telegram-project/pkg/http/handler"
	"github.com/mickey-mickser/telegram-project/pkg/storage/sheets"
	"github.com/sirupsen/logrus"
	_ "google.golang.org/api/sheets/v4"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func ctxWithSig() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	go func() {
		select {
		case <-ch:
			cancel()
		}

	}()

	return ctx, cancel
}
func main() {
	cfg, err := config.NewConfig("./configs/config.json")

	if err != nil {
		panic(err)
	}

	log := cfg.Log()

	ctx, cancel := ctxWithSig()
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			cancel()
		}
	}()

	cli := sheets.NewSheets(ctx).CliSheets()

	sheetUse := usecase.NewSheetUse(cli)

	wg := new(sync.WaitGroup)
	handlers := handler.NewHandler(log, sheetUse)
	http.NewHttp(log, handlers.InitRoutes()).Run(ctx, wg)

	logrus.Print("Server is running. Press CTRL+C to stop...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Shutting down server...")
	wg.Wait()
}
